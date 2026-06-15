package urll

import (
	"fmt"

	"github.com/kuche1/imotbg/config"
)

const _Part1 = "https://www.imot.bg/obiavi/prodazhbi/grad-sofiya/dvustaen"

// n 1-СТАЕН
// y 2-СТАЕН
// y 3-СТАЕН
// y 4-СТАЕН
// y МНОГОСТАЕН
// y МЕЗОНЕТ
// n АТЕЛИЕ, ТАВАН
const _Part2 = "p-%v?type_home=3~4~5~6~&price_min=%v&price_max=%v"

// y Завършени - Имоти въведени в експлоатация
// n В строеж или имоти НЕ въведени в експлоатация
const _Part3 = "&ybuild_type=1~"

func Generate(pageNum int, priceMinEur int, priceMaxEur int) string {
	url := _Part1

	if config.Garaj {
		url += "/s-garazh"
	}

	url += "/"
	url += fmt.Sprintf(_Part2, pageNum, priceMinEur, priceMaxEur)

	if config.Zavurshen {
		url += _Part3
	}

	return url
}
