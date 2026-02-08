package entities

type Persona struct {
	Nombre string
	Apellido string
	Edad   int
}

func (p Persona) Saludar() string {
	return "Hola, mi nombre es " + p.Nombre + " " + p.Apellido + " y tengo " + string(p.Edad) + " años."
}

func (p *Persona) CumplirAnios() {
	(*p).Edad++
}


