package house

import (
	"fmt"

	"github.com/dustin/go-humanize"
)

type House struct {
	Link             string
	PriceEur         float64
	Location         string
	AreaM2           int64
	Stai             string
	Ekstri           []string
	AdditionalParams map[string]string
}

func NewHouse(
	link string,
	priceEur float64,
	location string,
	areaM2 int64,
	stai string,
	ekstri []string,
	additionalParams map[string]string,
) *House {
	return &House{
		Link:             link,
		PriceEur:         priceEur,
		Location:         location,
		AreaM2:           areaM2,
		Stai:             stai,
		Ekstri:           ekstri,
		AdditionalParams: additionalParams,
	}
}

func (self *House) Sprintf() string {
	ekstri := fmt.Sprintf("%q", self.Ekstri)
	ekstri = ekstri[1 : len(ekstri)-1]

	additionalParams := fmt.Sprintf("%q", self.AdditionalParams)
	additionalParams = additionalParams[4 : len(additionalParams)-1]

	return fmt.Sprintf(
		`House:
    link             : %v
    price            : %v eur
    location         : %v
    area             : %v m2
    rooms            : %v
    ekstri           : %v
    additional params: %v
`,
		self.Link,
		humanize.Commaf(self.PriceEur),
		self.Location,
		self.AreaM2,
		self.Stai,
		ekstri,
		additionalParams,
	)
}
