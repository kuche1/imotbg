package extract

import (
	"log"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/kuche1/gonet"
	"github.com/kuche1/imotbg/config"
	"github.com/kuche1/imotbg/urll"
)

func extractSearchPages(net *gonet.Net, results chan<- *goquery.Document) {
	defer close(results)

	var wg sync.WaitGroup

	priceMax := config.PriceMax

	for priceMax >= config.PriceMin {
		priceMin := max(priceMax-config.PriceStep, config.PriceMin)

		// otherwise we are going to capture references ot the variables, and by the time
		// the thread has started the values will have changed
		anonPriceMin := priceMin
		anonPriceMax := priceMax

		wg.Go(func() {
			extractSearchPagesWithinPriceRange(net, results, anonPriceMin, anonPriceMax)
		})

		priceMax = priceMin - 1
	}

	wg.Wait()
}

func extractSearchPagesWithinPriceRange(
	net *gonet.Net,
	results chan<- *goquery.Document,
	priceMin int,
	priceMax int,
) {
	var pageNum int

	for pageNum = 1; pageNum <= config.LastPagePossible; pageNum++ {
		url := urll.Generate(pageNum, priceMin, priceMax)

		rawRespBytes := net.Req(url)

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(rawRespBytes)))
		if err != nil {
			log.Fatal(err)
		}

		// find the "Няма намерени обяви!" message
		if doc.Find("div.width980px.pageMessageAlert").Length() > 0 {
			break
		}

		results <- doc
	}

	if pageNum >= config.LastPagePossible {
		log.Fatal("The very last search page was reached, you need to reduce the price step")
		// log.Printf("The very last search page was reached, it is probable that some listings were omitted, you need to reduce the price step")
	}
}
