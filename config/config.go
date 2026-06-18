package config

import "math"

type Config struct {
	ZaduljitelnoGotovZaNanasqne bool
	ZaduljitelnoGaraj           bool
	PriceMinEur                 int
	PriceMaxEur                 int
	PloshtMinM2                 int64
	PloshtMaxM2                 int64
	StaiOkMap                   map[string]bool
}

func NewConfig() *Config {
	return &Config{
		ZaduljitelnoGotovZaNanasqne: true,
		ZaduljitelnoGaraj:           false,

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
	}
}
