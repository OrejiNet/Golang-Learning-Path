package gentleman

import "fmt"

func basic_loops() {

	// Defer es una función que se ejecuta al final de la función actual, justo antes de que esta retorne.
	defer fmt.Println("FIN")

	// Bucle Clásico
	for i := 0; i < 10; i++ {
		println("Iteración:", i)
	}

	// Bucle tipo while
	n := 0
	for n < 5 {
		println("Valor de n:", n)
		n++
	}

	// Bucle infinito
	for {
		n++
		if n == 5 {
			continue
		}

		fmt.Println("Valor de n en bucle infinito:", n)
		if n >= 10 {
			break
		}

	}

	// Range
	slice := []string {"Go", "Python", "JavaScript"}
	for index, value := range slice {
		fmt.Printf("Índice: %d, Valor: %s\n", index, value)
	}

	type  Digimon struct {
		Nombre string
	}

	digimons := []Digimon {
		{Nombre: "Agumon"},
		{Nombre: "Gabumon"},
		{Nombre: "Pomumon"},
		{Nombre: "Tentomon"},
		{Nombre: "Biyomon"},
		{Nombre: "Panimon"},
		{Nombre: "Otamamon"},
		{Nombre: "Kabuterimon"},
		{Nombre: "Agumon"},
		{Nombre: "Gabumon"},
		{Nombre: "Pomumon"},
		{Nombre: "Tentomon"},
		{Nombre: "Biyomon"},
		{Nombre: "Panimon"},
		{Nombre: "Otamamon"},
		{Nombre: "Kabuterimon"},
	}
	for _, digimon := range digimons {
		fmt.Println("Digimon:", digimon.Nombre)
	}

	type Tamer struct {
		Nombre string
		Digimons []Digimon
	}

	tamers := []Tamer {
		{
			Nombre: "Tai",	
			Digimons: []Digimon {
				{Nombre: "Agumon"},
				{Nombre: "Gabumon"},
			},
		},
		{
			Nombre: "Matt",	
			Digimons: []Digimon {
				{Nombre: "Pomumon"},
				{Nombre: "Tentomon"},
			},
		},
	}
	for _, tamer := range tamers {
		fmt.Println("Tamer:", tamer.Nombre)
		for _, digimon := range tamer.Digimons {
			fmt.Println("\tDigimon:", digimon.Nombre)
		}
	}
}
