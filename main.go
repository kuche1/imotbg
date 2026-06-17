package main

import (
	"fmt"
	"slices"

	"github.com/kuche1/imotbg/config"
	"github.com/kuche1/imotbg/extract"
	"github.com/kuche1/imotbg/house"
)

func main() {
	conf := config.NewConfig()

	houses := make([]*house.House, 0)

	for house := range extract.Main(conf) {
		houses = append(houses, house)
	}

	slices.SortFunc(
		houses,
		func(a *house.House, b *house.House) int {
			if a.AreaM2 < b.AreaM2 {
				return -1
			}
			if a.AreaM2 == b.AreaM2 {
				return 0
			}
			return 1
		},
	)

	for _, house := range houses {
		repr := house.Sprintf()
		fmt.Printf("%v\n", repr)
	}

	// TODO: add average price
}
