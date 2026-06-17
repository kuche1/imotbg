package config

type Config struct {
	ZaduljitelnoGotovZaNanasqne bool
	EdnostaenOk                 bool
	PriceMinEur                 int
	PriceMaxEur                 int
}

func NewConfig() *Config {
	return &Config{
		ZaduljitelnoGotovZaNanasqne: true,
		EdnostaenOk:                 true,

		PriceMinEur: 90_000,
		PriceMaxEur: 160_000,
	}
}
