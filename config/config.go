package config

import "math"

type Config struct {
	ZaduljitelnoGotovZaNanasqne bool
	PriceMinEur                 int
	PriceMaxEur                 int
	PloshtMinM2                 int64
	PloshtMaxM2                 int64
	StaiOkMap                   map[string]bool
	ZaduljitelniEkstri          []string
}

func NewConfig() *Config {
	return &Config{
		ZaduljitelnoGotovZaNanasqne: true,

		PriceMinEur: 0,
		PriceMaxEur: 190_000,

		PloshtMinM2: 0,
		PloshtMaxM2: math.MaxInt64,

		StaiOkMap: map[string]bool{
			"1-СТАЕН":       true,
			"2-СТАЕН":       true,
			"3-СТАЕН":       true,
			"4-СТАЕН":       true,
			"МНОГОСТАЕН":    true,
			"МЕЗОНЕТ":       true,
			"АТЕЛИЕ, ТАВАН": false,
		},

		// sekciq "Особености"
		ZaduljitelniEkstri: []string{
			"С гараж", "С паркинг",
		},
	}
}
