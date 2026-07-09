package config

import (
	"bufio"
	"log"
	"math"
	"os"
)

type Config struct {
	GotovZaNanasqne            bool
	PriceMinEur                int
	PriceMaxEur                int
	PloshtMinM2                int64
	PloshtMaxM2                int64
	StroitelstvoMissingOk      bool
	GodinaMissingOk            bool
	GodinaPredi1920Ok          bool
	GodinaMin                  int64
	GodinaMax                  int64
	StaiOkMap                  map[string]bool
	PoneEdnaZaduljitelnaEkstra []string
	AlreadyRegistered          []string
}

func NewConfig() *Config {
	return &Config{
		GotovZaNanasqne: true,

		PriceMinEur: 100_000,
		PriceMaxEur: 240_000,

		PloshtMinM2: 54,            //83,             //65,
		PloshtMaxM2: math.MaxInt64, //83,//65,

		StroitelstvoMissingOk: true,

		GodinaMissingOk:   true,
		GodinaPredi1920Ok: false,
		GodinaMin:         1960,
		GodinaMax:         2026,

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
		PoneEdnaZaduljitelnaEkstra: []string{
			// "С гараж", "С паркинг",
			"Тухла", // "ЕПК", "ПК",
		},

		AlreadyRegistered: parseAlreadyRegistered(),
	}
}

func parseAlreadyRegistered() []string {
	alreadyRegistered := make([]string, 0, 16)

	file, err := os.Open("results/registered/idk.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		if line[0] == '#' {
			continue
		}

		alreadyRegistered = append(alreadyRegistered, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return alreadyRegistered
}
