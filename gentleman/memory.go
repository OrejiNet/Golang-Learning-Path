package gentleman

import "fmt"

// punteros y memoria
// * -> puntero
// & -> dirección de memoria

//Numero es una copia del numero original, no afecta al valor original
func Incrementar(numero int){
	numero++
}

func IncrementarPuntero(numero *int){ //  <- * declara tipo "puntero a int"
	fmt.Println("Valor dentro de IncrementarPuntero:", *numero)
	fmt.Println("Dirección de memoria dentro de IncrementarPuntero:", numero)
	*numero++ // <- * desreferencia el puntero para acceder al valor original
}


func MostraPuntero(){
	puntero := new(int) // Crea un nuevo puntero a int
	fmt.Println("Valor del puntero:", *puntero)
	fmt.Println("Dirección de memoria del puntero:", puntero)
	*puntero = 10 // Asigna un valor al puntero
	fmt.Println("Valor del puntero después de asignar 10:", *puntero)
	//puntero = 10 // Asigna un nuevo valor al puntero, ahora apunta a la dirección de memoria de 10
	//fmt.Println("Valor del puntero después de asignar 10 directamente:", *puntero)
	//fmt.Println("Dirección de memoria del puntero después de asignar 10 directamente:", puntero)
}


// Paso 1 : Crear una variable de tipo int y asignarle un valor
// Paso 2 : Pasamos la direccion de memoria de la variable con ampersand
// Paso 3 : La función Incrementar recibe un puntero usando el asterisco
// Paso 4 : Incrementamos el valor al que apunta el puntero usando el asterisco 

func TestPunteros(){
	numero := 5
	Incrementar(numero)
	println("Valor después de Incrementar:", numero) // Imprime 5, no cambia el valor original
	
    IncrementarPuntero(&numero)  // & obtiene la dirección de numero
	println("Valor después de IncrementarPuntero:", numero) // Imprime 6, el valor original se incrementa
    // Dentro de la función:
    // numero (sin *) = 0xc000018030 (dirección de memoria)
    // *numero (con *) = 5 (el valor almacenado en esa dirección)
}