package extract

import (
	"log"
	"slices"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/kuche1/imotbg/config"
	"github.com/kuche1/imotbg/define"
)

const _ListingLinkPrefix = "https:"

func extractListingLinks(conf *config.Config, searchPages <-chan *goquery.Document, listingLinks chan<- string) {
	defer close(listingLinks)

	var wg sync.WaitGroup

	for range define.ThreadsExtractListingLinks {
		wg.Go(func() {
			extractListingLinksThr(conf, searchPages, listingLinks)
		})
	}

	wg.Wait()
}

func extractListingLinksThr(conf *config.Config, searchPages <-chan *goquery.Document, listingLinks chan<- string) {
	for doc := range searchPages {
		doc.Find("a.title.saveSlink").Each(func(i int, s *goquery.Selection) {
			href, exists := s.Attr("href")

			if !exists {
				log.Print("Could not find href\n")
				return
			}

			href = _ListingLinkPrefix + href

			if conf.ZaduljitelnoGotovZaNanasqne {
				if slices.Contains(define.BlacklistNeGotoviZaNanasqne, href) {
					return
				}
			}

			if slices.Contains(define.LinkBlacklist, href) {
				return
			}

			listingLinks <- href
		})
	}
}
