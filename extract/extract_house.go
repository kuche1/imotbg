package extract

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/kuche1/imotbg/config"
	"github.com/kuche1/imotbg/define"
	"github.com/kuche1/imotbg/house"
)

func extractHouses(conf *config.Config, listingLinks chan *_ListingPageData, houses chan *house.House) {
	defer close(houses)

	// IMPROVE: we can actually multithread this

	for pageData := range listingLinks {
		elemInfo := pageData.doc.Find("div.contactsBox").First()
		// elemParams := page_data.doc.Find("div.mainCarParams").First()

		if elemInfo.Length() == 0 {
			// log.Printf("Not found: %v", pageData.link)

			// err := os.WriteFile("debug", []byte(pageData.doc.Text()), 0644)
			// if err != nil {
			// 	panic(err)
			// }

			log.Fatalf("Could not find \"contacts box\": %v", pageData.link)

			// continue
		}

		// log.Printf("DBG: elemInfo=%v", elemShortInfo)
		// log.Printf("DBG: elemParams=%v", elemParams)

		price, priceInfo, invalid := findPrice(elemInfo, pageData.link)
		if invalid {
			continue
		}

		location, stai, invalid := findLocation(conf, elemInfo, pageData.link)
		if invalid {
			continue
		}

		elemParams := pageData.doc.Find("div.adParams").First()
		if elemParams.Length() == 0 {
			log.Fatalf("Could not find params: %v", pageData.link)
		}

		additionalParams,
			area,
			stroitelstvo,
			godina,
			invalid := findParams(
			conf,
			elemParams,
			pageData.link,
		)
		if invalid {
			continue
		}

		elemEkstri := pageData.doc.Find("div.carExtri").First()
		if elemEkstri.Length() == 0 {
			log.Fatalf("Could not find ekstri: %v", pageData.link)
		}

		ekstri, invalid := findEkstri(conf, elemEkstri, pageData.link)
		if invalid {
			continue
		}

		houses <- house.NewHouse(
			pageData.link,
			price,
			priceInfo,
			location,
			area,
			stai,
			stroitelstvo,
			godina,
			ekstri,
			additionalParams,
		)
	}
}

func findPrice(elemInfo *goquery.Selection, url string) (_price float64, _priceInfo string, _invalid bool) {
	const eur = "€"

	elem := elemInfo.Find("div.Price")
	if elem.Length() == 0 {
		log.Fatalf("Site must have changed layout - %v", url)
	}
	elem = elem.First()

	priceStr := strings.TrimSpace(elem.Text())

	if !strings.Contains(priceStr, eur) {
		log.Printf("Price in euros not found: %v", url)
		return 0, "", true
	}

	parts := strings.Split(priceStr, eur)

	priceStr = strings.TrimSpace(parts[0])
	priceStr = strings.ReplaceAll(priceStr, " ", "")

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		log.Fatalf("%v - %v", err, url)
	}

	//// actually im not sure if this really is the case I might be mistaken
	// // the shitty website sometimes shows you listings that are outside the price range
	// // that you have specified
	// if price > config.PriceMax {
	// 	return 0, true
	// }

	///// dds

	elem = elemInfo.Find("div.PriceInfo")
	if elem.Length() == 0 {
		log.Fatalf("Site must have changed layout - %v", url)
	}
	elem = elem.First()

	priceInfo := elem.Text()

	switch priceInfo {
	case "Не се начислява ДДС":
		fallthrough
	case "Цената е с включено ДДС":
		priceInfo = ""
	case "Цената е без ДДС":
		price *= 1.2 // vuzmojno e zaradi tova cenata da nadmine tazi setnata v konfiga
		priceInfo = fmt.Sprintf(" -> Добавени 20%% към цената заради: %v", priceInfo)
	default:
		log.Fatalf("Unrecognised DDS format `%v` - %v", priceInfo, url)
	}

	/////

	return price, priceInfo, false
}

func findLocation(conf *config.Config, elemInfo *goquery.Selection, link string) (_location string, _stai string, _blacklisted bool) {
	elemTitle := elemInfo.Find("div.obTitle").First()

	location := strings.TrimSpace(elemTitle.Text())
	// parts := strings.Split(location, "    ")
	// location = parts[1]
	// location = strings.TrimSpace(location)

	parts := strings.Split(location, "\n")
	if len(parts) != 4 {
		log.Fatalf("Could not parse location: %v", link)
	}
	// parts[0] -> Продава 3-СТАЕН
	// parts[1] ->     град София, Овча купел 2
	// parts[2] ->     ул. Светла (not always present)
	// parts[3] ->     Обява: 1c171137181218748

	///// location

	location = strings.TrimSpace(parts[1])

	// title_lower := strings.ToLower(title)

	if len(define.LocationPrefixWhitelist) > 0 {
		found := false

		for _, whitelisted := range define.LocationPrefixWhitelist {
			// whitelisted_title_lower := strings.ToLower(whitelisted_title)
			if strings.HasPrefix(location, whitelisted) {
				found = true
				break
			}
		}

		if !found {
			return "", "", true
		}
	}

	for _, blacklisted := range define.LocationPrefixBlacklist {
		// blacklisted := strings.ToLower(blacklisted)
		if strings.HasPrefix(location, blacklisted) {
			return "", "", true
		}
	}

	///// stai

	stai := parts[0]

	pref := "Продава "
	if !strings.HasPrefix(stai, pref) {
		log.Fatalf("Could not find prefix `%v` in `%v`: %v", pref, stai, link)
	}

	stai = stai[len(pref):]

	allowed, found := conf.StaiOkMap[stai]
	if !found {
		log.Fatalf("Number of rooms `%v` not found in room map `%v` - %v", stai, conf.StaiOkMap, link)
	}
	if !allowed {
		return "", "", true
	}

	/////

	return location, stai, false
}

func findParams(
	conf *config.Config,
	elemParams *goquery.Selection,
	link string,
) (
	_additionalParams map[string]string,
	_area int64,
	_stroitelstvo string,
	_godina int64,
	_blacklisted bool,
) {
	divs := elemParams.Find("div")
	if divs.Length() == 0 {
		log.Fatalf("unexpected layout")
	}

	parametri := make(map[string]string, divs.Length())

	divs.Each(func(i int, elem *goquery.Selection) {
		text := elem.Text()

		idx := strings.IndexByte(text, ':')
		if idx < 0 {
			log.Fatalf("Could not find separator in `%v` - %v", text, link)
		}

		key := text[:idx]
		value := text[idx+1:]

		_, existed := parametri[key]
		if existed {
			log.Fatalf("Duplicate param `%v` - %v", key, link)
		}

		parametri[key] = value
	})

	///// area

	areaStr := parametri["Площ"]
	delete(parametri, "Площ")

	suffix := " m2"
	if !strings.HasSuffix(areaStr, suffix) {
		log.Fatalf("Unexpected area format `%v` - %v", areaStr, link)
	}
	areaStr = strings.TrimSuffix(areaStr, suffix)

	area, err := strconv.ParseInt(areaStr, 10, 64)
	if err != nil {
		log.Fatalf("Area not a number for `%v` - %v", link, err)
	}

	if area < conf.PloshtMinM2 {
		return nil, 0, "", 0, true
	}

	if area > conf.PloshtMaxM2 {
		return nil, 0, "", 0, true
	}

	///// stroitelstvo, godina

	key := "Строителство"
	rawStroitelstvo, rawStroitelstvoOk := parametri[key]
	delete(parametri, key)

	stroitelstvo, godina, skip := findStroitelstvoGodina(conf, link, rawStroitelstvo, rawStroitelstvoOk)
	if skip {
		goto skip_house
	}

	///// return

	return parametri, area, stroitelstvo, godina, false

	///// skip

skip_house:
	return nil, 0, "", 0, true
}

func findStroitelstvoGodina(
	conf *config.Config,
	link string,
	rawStroitelstvo string,
	rawStroitelstvoOk bool,
) (
	_stroitelstvo string,
	_godina int64,
	_skip bool,
) {
	godina := int64(0)
	stroitelstvo := "<no data>"

	if !rawStroitelstvoOk {
		if conf.StroitelstvoMissingOk {
			goto return_data
		}
		goto skip_house
	}

	{
		t := ", "
		idx := strings.Index(rawStroitelstvo, t)
		if idx < 0 {
			log.Fatalf("Site layout must have changed - %v", link)
		}

		stroitelstvo = rawStroitelstvo[:idx]
		vuvedeno := rawStroitelstvo[idx+len(t):]

		// TODO: tozi kod e malko spaghett
		// trqbva year extraction-a da si e sobstvena funkciq

		if (vuvedeno == "") || (vuvedeno == "Въведен в експлоатация ") || (vuvedeno == "Не е въведен в експлоатация ") {
			// zachitam "Не е въведен в експлоатация" kato "missing godina" zashtoto
			// sum vidql pone 1 obqva, koqto e greshno vuvedeno tova
			if conf.GodinaMissingOk {
				goto return_data
			}
			goto skip_house
		}

		t = " г."
		if !strings.HasSuffix(vuvedeno, t) {
			log.Fatalf("Expected suffix `%v` for `%v` - %v", t, vuvedeno, link)
		}
		vuvedeno = vuvedeno[:len(vuvedeno)-len(t)]

		t = "Въведен в експлоатация "
		if strings.HasPrefix(vuvedeno, t) {
			vuvedeno = vuvedeno[len(t):]
		} else {
			t = "Ще бъде въведен в експлоатация "
			if strings.HasPrefix(vuvedeno, t) {
				// narochno ne proverqvam configa dali sum setnal "gotov za nanasqne"
				// v slu4ai 4e 1) nqkoi apartament e pochti gotov 2) nqkoi e missclick-nal
				// i e napisal greshno info za apartamenta
				vuvedeno = vuvedeno[len(t):]
			} else {

				val, err := strconv.ParseInt(vuvedeno, 10, 64)
				if err == nil {
					godina = val
					goto return_data
				}

				log.Fatalf("Unexpected prefix `%v` - %v", vuvedeno, link)
			}
		}

		if vuvedeno == "Преди 1920" {
			if conf.GodinaPredi1920Ok {
				goto return_data
			}
			goto skip_house
		}

		t = " - "
		idx = strings.Index(vuvedeno, t)
		if idx < 0 {

			val, err := strconv.ParseInt(vuvedeno, 10, 64)
			if err != nil {
				log.Fatalf("Year is not an integer `%v` - %v", vuvedeno, link)
			}
			godina = val

		} else {

			yearA := vuvedeno[:idx]
			yearB := vuvedeno[idx+len(t):]

			valA, err := strconv.ParseInt(yearA, 10, 64)
			if err != nil {
				log.Fatalf("Year is not an integer `%v` - %v", yearA, link)
			}

			valB, err := strconv.ParseInt(yearB, 10, 64)
			if err != nil {
				log.Fatalf("Year is not an integer `%v` - %v", yearB, link)
			}

			godina = (valA + valB) / 2

		}

	}

	if godina < conf.GodinaMin {
		goto skip_house
	}

return_data:
	return stroitelstvo, godina, false

skip_house:
	return "", 0, true
}

func findEkstri(conf *config.Config, elemEkstri *goquery.Selection, link string) (_value []string, _blacklisted bool) {
	elem := elemEkstri.Find("div.items").First()
	if elem.Length() == 0 {
		log.Fatalf("Could not find ekstri: %v", link)
	}

	raw := elem.Text()
	raw = strings.TrimSpace(raw)

	ekstriAvailable := strings.Split(raw, "\n")

	anyFound := false

	if len(conf.PoneEdnaZaduljitelnaEkstra) == 0 {
		anyFound = true
	} else {
		for _, ekstraRequired := range conf.PoneEdnaZaduljitelnaEkstra {
			if slices.Contains(ekstriAvailable, ekstraRequired) {
				anyFound = true
				break
			}
		}
	}

	if anyFound {
		return ekstriAvailable, false
	}

	return nil, true
}
