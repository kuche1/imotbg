package house

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/kuche1/imotbg/config"
)

type House struct {
	Link     string
	Price    float64
	Location string
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
	// title string,
	// engineType string,
	// horsepower int64,
	// yearProduced int16,
	// mialage int64,
	// gearbox string,
) *House {
	return &House{
		Link:     link,
		Price:    price,
		Location: location,
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
    price: %v %v
    location: %v
`,
		self.Link,
		humanize.Commaf(self.Price), config.Currency,
		self.Location,
		// self.EngineType,
		// self.Horsepower,
		// self.YearProduced,
		// self.Mialage,
		// self.Gearbox,
	)
}
