package main

import (
	"fmt"

	"github.com/kuche1/gonet"
	"github.com/kuche1/imotbg/config"
)

func main() {
	fmt.Printf("hi\n")

	net := gonet.NewNet(
		config.NetRequestDelayMS,
		config.NetCacheFolder,
		config.NetCachedResponseValiditySec,
	)

	fmt.Printf("net=%v\n", net)
}
