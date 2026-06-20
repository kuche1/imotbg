package define

import "runtime"

const NetCacheFolder = "net-cache"
const NetRequestDelayMS = 3
const NetCachedResponseValiditySec = 60 * 60 * 12 // 12 hours

const LastPagePossible = 25

var ThreadsExtractListingLinks = runtime.NumCPU()

const ExtractChanBuf = 64
