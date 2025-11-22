package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/lejkosir/VPSA-redovalnica/redovalnica"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "redovalnica",
		Usage: "Upravljanje ocen študentov",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "stOcen",
				Value: 2,
				Usage: "Minimalno število ocen za pozitivno oceno",
			},
			&cli.IntFlag{
				Name:  "minOcena",
				Value: 1,
				Usage: "Najmanjša možna ocena",
			},
			&cli.IntFlag{
				Name:  "maxOcena",
				Value: 10,
				Usage: "Največja možna ocena",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			studenti := map[string]redovalnica.Student{
				"63210001": {Ime: "Ana", Priimek: "Novak"},
				"63210002": {Ime: "Boris", Priimek: "Kralj"},
				"63210003": {Ime: "Janez", Priimek: "Zupan"},
			}

			redovalnica.DodajOceno(studenti, "63210001", 7)
			redovalnica.DodajOceno(studenti, "63210001", 8)
			redovalnica.DodajOceno(studenti, "63210001", 9)
			redovalnica.DodajOceno(studenti, "63210001", 10)
			redovalnica.DodajOceno(studenti, "63210002", 7)
			redovalnica.DodajOceno(studenti, "63210002", 6)
			redovalnica.DodajOceno(studenti, "63210003", 9)
			redovalnica.DodajOceno(studenti, "63210003", 8)
			redovalnica.DodajOceno(studenti, "63210003", 10)

			fmt.Println()
			redovalnica.IzpisVsehOcen(studenti)
			fmt.Println()
			redovalnica.IzpisiKoncniUspeh(studenti,
				cmd.Int("stOcen"),
				cmd.Int("minOcena"),
				cmd.Int("maxOcena"),
			)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
