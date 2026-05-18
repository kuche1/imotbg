package urll

import "fmt"

// n 1-СТАЕН
// y 2-СТАЕН
// y 3-СТАЕН
// y 4-СТАЕН
// y МНОГОСТАЕН
// y МЕЗОНЕТ
// n АТЕЛИЕ, ТАВАН
//
// up to 300_000 EUR
const _SprintfUrlPage = "https://www.imot.bg/obiavi/prodazhbi/grad-sofiya/dvustaen/p-%v?type_home=3~4~5~6~&price_max=%v"

func Generate(pageNum int, priceMaxEur int) string {
	return fmt.Sprintf(_SprintfUrlPage, pageNum, priceMaxEur)
}
