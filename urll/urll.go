package urll

import (
	"fmt"

	"github.com/kuche1/imotbg/config"
)

// TODO: this is suspicious, add a print for the finished URL

func Generate(conf *config.Config, pageNum int, priceMinEur int, priceMaxEur int) string {
	url := "https://www.imot.bg/obiavi/prodazhbi/grad-sofiya/ednostaen"

	///

	url += fmt.Sprintf("/p-%v", pageNum)

	///

	// y 1-СТАЕН
	// y 2-СТАЕН
	// y 3-СТАЕН
	// y 4-СТАЕН
	// y МНОГОСТАЕН
	// y МЕЗОНЕТ
	// y АТЕЛИЕ, ТАВАН
	url += "?type_home=2~3~4~5~6~8~"

	///

	url += fmt.Sprintf("&price_min=%v&price_max=%v", priceMinEur, priceMaxEur)

	///

	//// nqkoi obqvi burkat tozi tag
	// if conf.ZaduljitelnoGotovZaNanasqne {
	// 	// y Завършени - Имоти въведени в експлоатация
	// 	// n В строеж или имоти НЕ въведени в експлоатация
	// 	url += "&ybuild_type=1~"
	// }

	///

	return url
}
