package urll

import (
	"fmt"

	"github.com/kuche1/imotbg/define"
	"github.com/kuche1/imotbg/libconfig"
)

// TODO: this is suspicious, add a print for the finished URL

func Generate(config *libconfig.Config, pageNum int, priceMinEur int, priceMaxEur int) string {
	url := "https://www.imot.bg/obiavi/prodazhbi/grad-sofiya"

	///

	if define.EdnostaenOk {
		url += "/ednostaen"
	} else {
		url += "/dvustaen"
	}

	///

	if define.Garaj {
		url += "/s-garazh"
	}

	///

	url += fmt.Sprintf("/p-%v", pageNum)

	///

	url += "?type_home="

	if define.EdnostaenOk {
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

	///

	url += fmt.Sprintf("&price_min=%v&price_max=%v", priceMinEur, priceMaxEur)

	///

	if config.ZaduljitelnoGotovZaNanasqne {
		// y Завършени - Имоти въведени в експлоатация
		// n В строеж или имоти НЕ въведени в експлоатация
		url += "&ybuild_type=1~"
	}

	///

	return url
}
