package extract

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/kuche1/gonet"
	"github.com/kuche1/imotbg/config"
	"github.com/kuche1/imotbg/define"
	"github.com/kuche1/imotbg/house"
)

func Main(conf *config.Config) chan *house.House {
	net := gonet.NewNet(
		define.NetRequestDelayMS,
		define.NetCacheFolder,
		define.NetCachedResponseValiditySec,
	)

	searchPages := make(chan *goquery.Document, define.ExtractChanBuf)
	go extractSearchPages(conf, net, searchPages)

	listingLinks := make(chan string, define.ExtractChanBuf)
	go extractListingLinks(conf, searchPages, listingLinks)

	listingPageData := make(chan *_ListingPageData, define.ExtractChanBuf)
	go downloadListingPages(net, listingLinks, listingPageData)

	houses := make(chan *house.House, define.ExtractChanBuf)
	go extractHouses(conf, listingPageData, houses)

	return houses
}
