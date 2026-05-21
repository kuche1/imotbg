package house

import (
	"fmt"

	"github.com/kuche1/imotbg/config"
)

type House struct {
	Link  string
	Price float64
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
	// title string,
	// engineType string,
	// horsepower int64,
	// yearProduced int16,
	// mialage int64,
	// gearbox string,
) *House {
	return &House{
		Link:  link,
		Price: price,
		// Title:        title,
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
`,
		self.Link,
		self.Price, config.Currency,
		// self.Title,
		// self.EngineType,
		// self.Horsepower,
		// self.YearProduced,
		// self.Mialage,
		// self.Gearbox,
	)
}
