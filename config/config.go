// Favoriti:
// https://www.imot.bg/obiava-1c175369817671057-prodava-tristaen-apartament-grad-sofiya-gotse-delchev

package config

import "math"

type Config struct {
	ZaduljitelnoGotovZaNanasqne bool
	ZaduljitelnoGaraj           bool
	EdnostaenOk                 bool
	PriceMinEur                 int
	PriceMaxEur                 int
	PloshtMinM2                 int64
	PloshtMaxM2                 int64
}

func NewConfig() *Config {
	return &Config{
		ZaduljitelnoGotovZaNanasqne: true,
		ZaduljitelnoGaraj:           false,
		EdnostaenOk:                 true,

		PriceMinEur: 90_000,
		PriceMaxEur: 180_000,

		PloshtMinM2: 0,
		PloshtMaxM2: math.MaxInt64,
	}
}
