package libconfig

type Config struct {
	ZaduljitelnoGotovZaNanasqne bool
}

func NewConfig() *Config {
	return &Config{
		ZaduljitelnoGotovZaNanasqne: true,
	}
}
