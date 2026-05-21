package extract

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/kuche1/gonet"
	"github.com/kuche1/imotbg/config"
)

func Main() {
	net := gonet.NewNet(
		config.NetRequestDelayMS,
		config.NetCacheFolder,
		config.NetCachedResponseValiditySec,
	)

	searchPages := make(chan *goquery.Document, config.ExtractChanBuf)
	go extractSearchPages(net, searchPages)

	listingLinks := make(chan string, config.ExtractChanBuf)
	go extractListingLinks(searchPages, listingLinks)

	listingPageData := make(chan *_ListingPageData, config.ExtractChanBuf)
	go downloadListingPages(net, listingLinks, listingPageData)

	for data := range listingPageData {
		fmt.Printf("data=%v\n", data)
	}
}
