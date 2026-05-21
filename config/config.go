package config

import "runtime"

const NetCacheFolder = "net-cache"
const NetRequestDelayMS = 2                       // TODO: see if we can safely lower this value
const NetCachedResponseValiditySec = 60 * 60 * 10 // 10 hours

const LastPagePossible = 25

// in EUR
const PriceMin = 100_000
const PriceMax = 150_000
const PriceStep = 10_000 // if this is too big you may miss some listings

const Currency = "EUR"

var ThreadsExtractListingLinks = runtime.NumCPU()

const ListingLinkPrefix = "https:"

var LinkBlacklist = []string{
	// "https://www.mobile.bg/obiava-11767956291631060-bmw-740-li-xdrive",
}

const ExtractChanBuf = 64

var LocationBlacklist = []string{
	// "град София, Филиповци",
	// "град София, Модерно предградие",
	// "град София, с. Мало Бучино",
	// "град София, Център",
	// "град София, Надежда 3",
	// "град София, Овча купел 2",
}
