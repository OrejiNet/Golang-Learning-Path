package gentleman

import (
	"errors"
	"fmt"
)

func Dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("No se puede dividir por 0")
	}
	cociente := a / b
	return cociente, nil
}

func PrintNombres(nombres ...string){
	for _, nombre := range nombres{
		fmt.Println("Nombre:", nombre)
	}
}

func Contador() func() int {
	count := 0
	return func() int{
		count++
		return count
	}
}

func CallContador(times int){
	contador := Contador()
	for range times {
		fmt.Println(contador())
	}
}

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	for i := n - 1; i > 0; i-- {
		n *= i
	}
	return n
} 


// No existen clases, pero se pueden definir tipos personalizados con struct e Interfaces
type Rectangulo struct {
	Ancho, Alto float64
}

func (r Rectangulo) Area() float64 {
	return r.Ancho * r.Alto
}

func (r Rectangulo) Perimetro() float64 {
	return 2*(r.Ancho + r.Alto)
}


func InfoRectangulo(r Rectangulo) {
	fmt.Printf("Rectángulo: Ancho=%.2f, Alto=%.2f, Área=%.2f, Perímetro=%.2f\n", r.Ancho, r.Alto, r.Area(), r.Perimetro())
}

