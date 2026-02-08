package gentleman

import "fmt"

// Las interfaces no se implementan explícitamente, 
// cualquier tipo que tenga los métodos definidos en la interfaz se considera que implementa esa interfaz
type Forma interface{
	Area() float64
	Perimetro() float64
}

type Circulo struct{
	Radio float64
}

func (c Circulo) Area() float64{
	return 3.14 * c.Radio * c.Radio
}

func (c Circulo) Perimetro() float64{
	return 2 * 3.14 * c.Radio
}

func Info(r Forma){
	fmt.Println("Área:", r.Area())
	fmt.Println("Perímetro:", r.Perimetro())
}

func CallInfo(){
	circulo := Circulo{Radio: 5}
	Info(circulo)
}

