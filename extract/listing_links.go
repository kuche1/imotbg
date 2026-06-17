package extract

import (
	"log"
	"slices"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/kuche1/imotbg/config"
)

const _ListingLinkPrefix = "https:"

func extractListingLinks(searchPages <-chan *goquery.Document, listingLinks chan<- string) {
	defer close(listingLinks)

	var wg sync.WaitGroup

	for range config.ThreadsExtractListingLinks {
		wg.Go(func() {
			extractListingLinksThr(searchPages, listingLinks)
		})
	}

	wg.Wait()
}

func extractListingLinksThr(searchPages <-chan *goquery.Document, listingLinks chan<- string) {
	for doc := range searchPages {
		doc.Find("a.title.saveSlink").Each(func(i int, s *goquery.Selection) {
			href, exists := s.Attr("href")

			if !exists {
				log.Print("Could not find href\n")
				return
			}

			href = _ListingLinkPrefix + href

			if config.GotovZaNanasqne {
				if slices.Contains(config.BlacklistNeGotoviZaNanasqne, href) {
					return
				}
			}

			if slices.Contains(config.LinkBlacklist, href) {
				return
			}

			listingLinks <- href
		})
	}
}
