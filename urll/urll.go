package urll

import (
	"fmt"

	"github.com/kuche1/imotbg/config"
)

// TODO: this is suspicious, add a print for the finished URL

func Generate(conf *config.Config, pageNum int, priceMinEur int, priceMaxEur int) string {
	url := "https://www.imot.bg/obiavi/prodazhbi/grad-sofiya"

	///

	if conf.EdnostaenOk {
		url += "/ednostaen"
	} else {
		url += "/dvustaen"
	}

	///

	if conf.ZaduljitelnoGaraj {
		url += "/s-garazh"
	}

	///

	url += fmt.Sprintf("/p-%v", pageNum)

	///

	url += "?type_home="

	if conf.EdnostaenOk {
		// y 1-СТАЕН
		url += "2~"
	}

	// y 2-СТАЕН
	// y 3-СТАЕН
	// y 4-СТАЕН
	// y МНОГОСТАЕН
	// y МЕЗОНЕТ
	// n АТЕЛИЕ, ТАВАН
	url += "3~4~5~6~"
	// TODO: moje bi vmesto da si igraq s URL-a, prosto
	// da parse-vam dali e 1-staen, 2-staen i tn

	///

	url += fmt.Sprintf("&price_min=%v&price_max=%v", priceMinEur, priceMaxEur)

	///

	if conf.ZaduljitelnoGotovZaNanasqne {
		// y Завършени - Имоти въведени в експлоатация
		// n В строеж или имоти НЕ въведени в експлоатация
		url += "&ybuild_type=1~"
	}

	///

	return url
}
