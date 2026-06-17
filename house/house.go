package house

import (
	"fmt"

	"github.com/dustin/go-humanize"
)

type House struct {
	Link     string
	PriceEur float64
	Location string
	AreaM2   int64
	// Title        string
	// EngineType   string
	// Horsepower   int64
	// YearProduced int16
	// Mialage      int64
	// Gearbox      string
}

func NewHouse(
	link string,
	price float64,
	location string,
	area int64,
	// title string,
	// engineType string,
	// horsepower int64,
	// yearProduced int16,
	// mialage int64,
	// gearbox string,
) *House {
	return &House{
		Link:     link,
		PriceEur: price,
		Location: location,
		AreaM2:   area,
		// EngineType:   engineType,
		// Horsepower:   horsepower,
		// YearProduced: yearProduced,
		// Mialage:      mialage,
		// Gearbox:      gearbox,
	}
}

func (self *House) Sprintf() string {
	return fmt.Sprintf(
		`House:
    link: %v
    price: %v [EUR]
    location: %v
    area: %v [m2]
`,
		self.Link,
		humanize.Commaf(self.PriceEur),
		self.Location,
		self.AreaM2,
		// self.EngineType,
		// self.Horsepower,
		// self.YearProduced,
		// self.Mialage,
		// self.Gearbox,
	)
}
