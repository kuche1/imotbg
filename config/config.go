package config

import "math"

type Config struct {
	ZaduljitelnoGotovZaNanasqne bool
	PriceMinEur                 int
	PriceMaxEur                 int
	PloshtMinM2                 int64
	PloshtMaxM2                 int64
	StroitelstvoMissingOk       bool
	StaiOkMap                   map[string]bool
	PoneEdnaZaduljitelnaEkstra  []string
}

func NewConfig() *Config {
	return &Config{
		ZaduljitelnoGotovZaNanasqne: true,

		PriceMinEur: 0,
		PriceMaxEur: 200_000,

		PloshtMinM2: 40,            //65,
		PloshtMaxM2: math.MaxInt64, //65,

		StaiOkMap: map[string]bool{
			"1-СТАЕН":       true,
			"2-СТАЕН":       true,
			"3-СТАЕН":       true,
			"4-СТАЕН":       true,
			"МНОГОСТАЕН":    true,
			"МЕЗОНЕТ":       true,
			"АТЕЛИЕ, ТАВАН": false,
		},

		StroitelstvoMissingOk: true,

		// sekciq "Особености"
		PoneEdnaZaduljitelnaEkstra: []string{
			// "С гараж", "С паркинг",
			// "Тухла", "ЕПК", "ПК",
		},
	}
}
