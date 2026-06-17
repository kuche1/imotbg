package extract

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/kuche1/gonet"
	"github.com/kuche1/imotbg/define"
	"github.com/kuche1/imotbg/house"
	"github.com/kuche1/imotbg/libconfig"
)

func Main(config *libconfig.Config) chan *house.House {
	net := gonet.NewNet(
		define.NetRequestDelayMS,
		define.NetCacheFolder,
		define.NetCachedResponseValiditySec,
	)

	searchPages := make(chan *goquery.Document, define.ExtractChanBuf)
	go extractSearchPages(config, net, searchPages)

	listingLinks := make(chan string, define.ExtractChanBuf)
	go extractListingLinks(config, searchPages, listingLinks)

	listingPageData := make(chan *_ListingPageData, define.ExtractChanBuf)
	go downloadListingPages(net, listingLinks, listingPageData)

	houses := make(chan *house.House, define.ExtractChanBuf)
	go extractHouses(listingPageData, houses)

	return houses
}
