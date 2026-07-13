package config

import (
	"bufio"
	"log"
	"math"
	"os"
	"strings"
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
		PriceMaxEur: 250_000, // 240_000

		PloshtMinM2: 54,            //83,             //65,
		PloshtMaxM2: math.MaxInt64, //83,//65,

		StroitelstvoMissingOk: true,

		GodinaMissingOk:   true,
		GodinaPredi1920Ok: false,
		GodinaMin:         1980,
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

	folderPath := "results/registered"

	entries, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filePath := folderPath + "/" + entry.Name()

		fileObj, err := os.Open(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer fileObj.Close()

		scanner := bufio.NewScanner(fileObj)
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) == 0 {
				continue
			}

			if line[0] == '#' {
				continue
			}

			if !strings.HasPrefix(line, "https://") {
				log.Fatalf("Invalid line in config (must be a link or a comment): %v", line)
			}

			alreadyRegistered = append(alreadyRegistered, line)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	return alreadyRegistered
}
