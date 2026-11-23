package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/JonasJavornik/vremenska-postaja/vreme"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "vremenska-postaja",
		Usage: "Beleženje temperaturnih meritev vremenskih postaj in izračun statistike",
		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:  "minTemp",
				Value: -50,
				Usage: "najnižja dovoljena temperatura meritve (°C)",
			},
			&cli.Float64Flag{
				Name:  "maxTemp",
				Value: 50,
				Usage: "najvišja dovoljena temperatura meritve (°C)",
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Value: 3,
				Usage: "najmanjše število meritev, potrebnih za izračun statistike",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			minTemp := cmd.Float64("minTemp")
			maxTemp := cmd.Float64("maxTemp")
			stMeritev := cmd.Int("stMeritev")

			postaje := make(map[string]vreme.Postaja)

			// Ustvarjanje postaj preko izvožene funkcije DodajPostajo
			vreme.DodajPostajo(postaje, "KP01", "Luka Koper", "Koper")
			vreme.DodajPostajo(postaje, "KR01", "Kredarica", "Triglav")
			vreme.DodajPostajo(postaje, "LJ01", "Ljubljana Bežigrad", "Ljubljana")
			vreme.DodajPostajo(postaje, "MB01", "Maribor Tabor", "Maribor")

			// Dodajanje začetnih meritev
			for _, t := range []float64{24.5, 27, 26.1, 28.4} {
				vreme.DodajMeritev(postaje, "KP01", t, minTemp, maxTemp)
			}
			for _, t := range []float64{2, -1.5, 3.5} {
				vreme.DodajMeritev(postaje, "KR01", t, minTemp, maxTemp)
			}
			for _, t := range []float64{18.2, 21.5, 19.8, 16.5} {
				vreme.DodajMeritev(postaje, "LJ01", t, minTemp, maxTemp)
			}
			vreme.DodajMeritev(postaje, "MB01", 12.4, minTemp, maxTemp)

			// Izpis meritev in poročila
			vreme.IzpisMeritev(postaje)
			fmt.Println()
			vreme.IzpisPorocila(postaje, stMeritev)
			fmt.Println()

			// Demonstracija robnih primerov (neveljavna temperatura in neobstoječa postaja)
			fmt.Println("PREVERJANJE ROBNIH PRIMEROV:")
			vreme.DodajMeritev(postaje, "LJ01", 60.0, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "CE01", 15.0, minTemp, maxTemp)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
