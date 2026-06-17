package main

import (
	"fmt"

	"github.com/kuche1/imotbg/config"
	"github.com/kuche1/imotbg/extract"
)

func main() {
	conf := config.NewConfig()

	houses := extract.Main(conf)

	for house := range houses {
		repr := house.Sprintf()
		fmt.Printf("%v\n", repr)
	}

	// TODO: add average price
}
