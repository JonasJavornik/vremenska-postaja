// Paket vreme omogoča upravljanje vremenskih postaj, beleženje temperaturnih meritev, izračun in izpis osnovne statistike meritev.
package vreme

import (
	"fmt"
	"sort"
)

// Postaja = posamezna vremensko postajo z imenom, krajem in seznamom zabeleženih temp. meritev.
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// DodajPostajo v slovar postaj doda novo vremensko postajo s podano oznako, imenom in krajem. Če postaja že obstaja, izpiše obvestilo.
func DodajPostajo(postaje map[string]Postaja, oznaka, ime, kraj string) {
	if _, obstaja := postaje[oznaka]; obstaja {
		fmt.Printf("Opozorilo: postaja z oznako %q že obstaja.\n", oznaka)
		return
	}

	postaje[oznaka] = Postaja{
		ime:  ime,
		kraj: kraj,
	}
}

// DodajMeritev postaji z dano oznako doda novo temp. meritev, če je temp. znotraj intervala [minTemp, maxTemp].
// Če postaja ne obstaja/temperatura izven območja, funkcija izpiše ustrezno obvestilo.
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura, minTemp, maxTemp float64) {
	if temperatura < minTemp || temperatura > maxTemp {
		fmt.Printf("Napaka: temperatura %.1f °C ni v dovoljenem območju (%.1f °C do %.1f °C).\n", temperatura, minTemp, maxTemp)
		return
	}

	postaja, obstaja := postaje[oznaka]
	if !obstaja {
		fmt.Printf("Napaka: postaja z oznako %q ne obstaja.\n", oznaka)
		return
	}

	postaja.meritve = append(postaja.meritve, temperatura)
	postaje[oznaka] = postaja
}

// statistika izračuna najnižjo, najvišjo in povprečno temperaturo izbrane postaje.
// Če postaja ne obstaja/ima manj kot stMeritev meritev, vrne ok == false.
func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	postaja, obstaja := postaje[oznaka]
	if !obstaja || len(postaja.meritve) < stMeritev {
		return 0, 0, 0, false
	}

	najnizja = postaja.meritve[0]
	najvisja = postaja.meritve[0]
	vsota := 0.0

	for _, t := range postaja.meritve {
		if t < najnizja {
			najnizja = t
		}
		if t > najvisja {
			najvisja = t
		}
		vsota += t
	}

	povprecna = vsota / float64(len(postaja.meritve))
	return najnizja, najvisja, povprecna, true
}

// IzpisMeritev urejeno po oznakah izpiše vse vremenske postaje, njihove kraje in vse zabeležene temperaturne meritve.
func IzpisMeritev(postaje map[string]Postaja) {
	oznake := urejeneOznake(postaje)

	fmt.Println("MERITVE:")
	for _, oznaka := range oznake {
		p := postaje[oznaka]
		fmt.Printf("%s - %s (%s): %v\n", oznaka, p.ime, p.kraj, p.meritve)
	}
}

// IzpisPorocila za vsako postajo izpiše najnižjo, najvišjo in povprečno temp. ter vremenski opis glede na povprečje.
// Če ima postaja manj kot stMeritev meritev, izpiše obvestilo o premajhnem številu meritev.
func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	oznake := urejeneOznake(postaje)

	for _, oznaka := range oznake {
		p := postaje[oznaka]
		min, max, povp, ok := statistika(postaje, oznaka, stMeritev)
		if !ok {
			fmt.Printf("%s: premalo meritev\n", p.ime)
			continue
		}

		var opis string
		switch {
		case povp >= 25:
			opis = "Vroče"
		case povp >= 15:
			opis = "Toplo"
		case povp >= 5:
			opis = "Hladno"
		default:
			opis = "Mrzlo"
		}

		fmt.Printf("%s: min %.1f °C, max %.1f °C, povprečje %.1f °C -> %s\n", p.ime, min, max, povp, opis)
	}
}

func urejeneOznake(postaje map[string]Postaja) []string {
	oznake := make([]string, 0, len(postaje))
	for oznaka := range postaje {
		oznake = append(oznake, oznaka)
	}
	sort.Strings(oznake)
	return oznake
}
