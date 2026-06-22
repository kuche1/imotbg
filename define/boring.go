package define

import "runtime"

const NetCacheFolder = "net-cache"
const NetRequestDelayMS = 3
const NetCachedResponseValiditySec = 60 * 60 * 12 // 12 hours

const LastPagePossible = 25

var ThreadsExtractListingLinks = runtime.NumCPU()

const ExtractChanBuf = 64

const PriceStepEur = 6_000 // if this is too big you may miss some listings
