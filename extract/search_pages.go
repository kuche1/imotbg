package extract

import (
	"log"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/kuche1/gonet"
	"github.com/kuche1/imotbg/config"
	"github.com/kuche1/imotbg/define"
	"github.com/kuche1/imotbg/urll"
)

func extractSearchPages(conf *config.Config, net *gonet.Net, results chan<- *goquery.Document) {
	defer close(results)

	var wg sync.WaitGroup

	priceMax := conf.PriceMaxEur

	for priceMax >= conf.PriceMinEur {
		priceMin := max(priceMax-define.PriceStepEur, conf.PriceMinEur)

		// otherwise we are going to capture references ot the variables, and by the time
		// the thread has started the values will have changed
		anonPriceMin := priceMin
		anonPriceMax := priceMax

		wg.Go(func() {
			extractSearchPagesWithinPriceRange(conf, net, results, anonPriceMin, anonPriceMax)
		})

		priceMax = priceMin - 1
	}

	wg.Wait()
}

func extractSearchPagesWithinPriceRange(
	conf *config.Config,
	net *gonet.Net,
	results chan<- *goquery.Document,
	priceMin int,
	priceMax int,
) {
	var pageNum int

	for pageNum = 1; pageNum <= define.LastPagePossible; pageNum++ {
		url := urll.Generate(conf, pageNum, priceMin, priceMax)

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

	if pageNum >= define.LastPagePossible {
		log.Fatal("The very last search page was reached, you need to reduce the price step")
	}
}
