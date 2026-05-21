package extract

import (
	"log"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/kuche1/gonet"
	"golang.org/x/net/html/charset"
)

type _ListingPageData struct {
	link string
	doc  *goquery.Document
}

func downloadListingPages(net *gonet.Net, listingLinks chan string, listingPageData chan<- *_ListingPageData) {
	defer close(listingPageData)

	var wg sync.WaitGroup

	for link := range listingLinks {
		wg.Go(func() {
			downloadListingPagesThread(net, listingPageData, link)
		})
	}

	wg.Wait()
}

func downloadListingPagesThread(net *gonet.Net, listingPageData chan<- *_ListingPageData, link string) {
	// fmt.Printf("downloadCarPages: got car page link: %v\n", link)

	page_bytes := net.Req(link)
	page_text := string(page_bytes)

	reader0 := strings.NewReader(page_text)

	reader1, err := charset.NewReaderLabel("windows-1251", reader0)
	if err != nil {
		panic(err)
	}

	doc, err := goquery.NewDocumentFromReader(reader1)
	if err != nil {
		log.Fatal(err)
	}

	listingPageData <- &_ListingPageData{
		link,
		doc,
	}
}
