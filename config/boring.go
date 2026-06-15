package config

import "runtime"

const NetCacheFolder = "net-cache"
const NetRequestDelayMS = 2                       // TODO: see if we can safely lower this value
const NetCachedResponseValiditySec = 60 * 60 * 10 // 10 hours

const LastPagePossible = 25

const Currency = "EUR"

var ThreadsExtractListingLinks = runtime.NumCPU()

const ExtractChanBuf = 64
