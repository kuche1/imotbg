package extract

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/kuche1/gonet"
	"github.com/kuche1/imotbg/config"
	"github.com/kuche1/imotbg/house"
)

func Main() chan *house.House {
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

	houses := make(chan *house.House, config.ExtractChanBuf)
	go extractHouses(listingPageData, houses)

	return houses
}
