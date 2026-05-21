package config

const NetCacheFolder = "net-cache"
const NetRequestDelayMS = 1
const NetCachedResponseValiditySec = 60 * 60 * 6 // 6 hours

const LastPagePossible = 25

// in EUR
const PriceMin = 100_000
const PriceMax = 150_000
const PriceStep = 10_000 // if this is too big you may miss some listings
