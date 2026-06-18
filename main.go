// Findings:
// https://www.imot.bg/obiava-1c175369817671057-prodava-tristaen-apartament-grad-sofiya-gotse-delchev
// https://www.imot.bg/obiava-1c178152298876317-prodava-tristaen-apartament-grad-sofiya-mladost-2#&gid=1&pid=2
// https://www.imot.bg/obiava-1c177712109058420-prodava-tristaen-apartament-grad-sofiya-mladost-1

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
		func(houseA *house.House, houseB *house.House) int {
			formula := func(house *house.House) float64 {
				return float64(house.AreaM2)
				//coeffArea := math.Pow(float64(house.AreaM2) / 64.0, 1.2) // float64(house.AreaM2)
				//coeffPrice := 1 / math.Pow(float64(house.PriceEur) / 150_000.0, 1.1)
				//return coeffArea * coeffPrice
			}

			a := formula(houseA)
			b := formula(houseB)

			if a < b {
				return -1
			}
			if a == b {
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
