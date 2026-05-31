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
	// za remont
	"https://www.imot.bg/obiava-1c173679076198229-prodava-tristaen-apartament-grad-sofiya-tsentar-bul-slivnitsa",
	"https://www.imot.bg/obiava-1c175005241220903-prodava-tristaen-apartament-grad-sofiya-tsentar",
	"https://www.imot.bg/obiava-1b177789507250802-prodava-dvustaen-apartament-grad-sofiya-tsentar-ul-sofroniy-vrachanski",
}

const ExtractChanBuf = 64

// active only if ccontains at lest 1 item
var LocationPrefixWhitelist = []string{
	"град София, Център",

	// "град София, Овча купел", // covers the base, 1 and 2
}

var LocationPrefixBlacklist = []string{
	// "град София, Обеля 2",
	// "град София, Суходол",
	// "град София, Люлин 9",
	// "град София, Модерно предградие",
	// "град София, Обеля",
	// "град София, Илинден",
	// "град София, Симеоново",
	// "град София, Фондови жилища",
	// "град София, Овча купел 1",
	// "град София, Овча купел 2",
	// "град София, Толстой",
	// "град София, Филиповци",
}
