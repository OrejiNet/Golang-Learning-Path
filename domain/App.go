package domain

import "curso-go/domain/entities"

func App() {
	persona := entities.Persona{
		Nombre:   "Cristopher",
		Apellido: "Angulo",
		Edad:     20,
	}

	persona.Saludar()
	persona.CumplirAnios()
	persona.Saludar()
}
