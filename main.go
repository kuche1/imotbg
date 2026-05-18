package main

import (
	"fmt"

	"github.com/kuche1/gonet"
	"github.com/kuche1/imotbg/config"
	"github.com/kuche1/imotbg/urll"
)

func main() {
	fmt.Printf("hi\n")

	net := gonet.NewNet(
		config.NetRequestDelayMS,
		config.NetCacheFolder,
		config.NetCachedResponseValiditySec,
	)

	fmt.Printf("net=%v\n", net)

	pageUrl := urll.Generate(1, 150_000)
	fmt.Printf("pageUrl=%v\n", pageUrl)
}
