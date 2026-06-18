package define

import "runtime"

const NetCacheFolder = "net-cache"
const NetRequestDelayMS = 3
const NetCachedResponseValiditySec = 60 * 60 * 120000 // 120000 hours // TODO: revert to 12

const LastPagePossible = 25

var ThreadsExtractListingLinks = runtime.NumCPU()

const ExtractChanBuf = 64
