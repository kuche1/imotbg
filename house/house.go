package house

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/kuche1/imotbg/define"
)

type House struct {
	Link     string
	Price    float64
	Location string
	Area     int64
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
		Price:    price,
		Location: location,
		Area:     area,
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
    area: %v
`,
		self.Link,
		humanize.Commaf(self.Price), define.Currency,
		self.Location,
		self.Area,
		// self.EngineType,
		// self.Horsepower,
		// self.YearProduced,
		// self.Mialage,
		// self.Gearbox,
	)
}
