package config

type Config struct {
	ZaduljitelnoGotovZaNanasqne bool
	PriceMinEur                 int
	PriceMaxEur                 int
}

func NewConfig() *Config {
	return &Config{
		ZaduljitelnoGotovZaNanasqne: true,
		PriceMinEur:                 90_000,
		PriceMaxEur:                 150_000,
	}
}
