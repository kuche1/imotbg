package main

import (
	"fmt"

	"github.com/kuche1/imotbg/extract"
	"github.com/kuche1/imotbg/libconfig"
)

func main() {
	config := libconfig.NewConfig()

	houses := extract.Main(config)

	for house := range houses {
		repr := house.Sprintf()
		fmt.Printf("%v\n", repr)
	}

	// TODO: add average price
}
