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

	slices.SortStableFunc(
		houses,
		func(houseA *house.House, houseB *house.House) int {
			formula := func(house *house.House) float64 {
				return float64(house.AreaM2)
			}
			a := formula(houseA)
			b := formula(houseB)
			if a < b {
				return -1
			}
			if b > 1 {
				return 1
			}
			return 0
		},
	)

	// slices.SortStableFunc(
	// 	houses,
	// 	func(houseA *house.House, houseB *house.House) int {
	// 		formula := func(house *house.House) float64 {
	// 			coeffPrice := 1.0 / float64(house.PriceEur)
	// 			coeffArea := math.Pow(float64(house.AreaM2), 1.3)
	// 			return coeffPrice * coeffArea

	// 			// return float64(house.AreaM2)

	// 			//coeffArea := math.Pow(float64(house.AreaM2) / 64.0, 1.2) // float64(house.AreaM2)
	// 			//coeffPrice := 1 / math.Pow(float64(house.PriceEur) / 150_000.0, 1.1)
	// 			//return coeffArea * coeffPrice
	// 		}

	// 		a := formula(houseA)
	// 		b := formula(houseB)

	// 		if a < b {
	// 			return -1
	// 		}
	// 		if b > 1 {
	// 			return 1
	// 		}
	// 		if houseA.Link < houseB.Link {
	// 			return -1
	// 		}
	// 		if houseA.Link > houseB.Link {
	// 			return 1
	// 		}
	// 		panic("this is still possible, if we get duplicate links")
	// 		return 0
	// 	},
	// )

	// houses = houses[:len(houses)-17]

	for _, house := range houses {
		repr := house.Sprintf()
		fmt.Printf("%v\n", repr)
	}

	// TODO: add average price (and calc aobzavedeni vs ne)
}
