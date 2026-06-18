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
	Stai     string
	Ekstri   []string
}

func NewHouse(
	link string,
	priceEur float64,
	location string,
	areaM2 int64,
	stai string,
	ekstri []string,
) *House {
	return &House{
		Link:     link,
		PriceEur: priceEur,
		Location: location,
		AreaM2:   areaM2,
		Stai:     stai,
		Ekstri:   ekstri,
	}
}

func (self *House) Sprintf() string {
	return fmt.Sprintf(
		`House:
    link    : %v
    price   : %v eur
    location: %v
    area    : %v m2
    rooms   : %v
    ekstri  : %q
`,
		self.Link,
		humanize.Commaf(self.PriceEur),
		self.Location,
		self.AreaM2,
		self.Stai,
		self.Ekstri,
	)
}
