package cmd

import "fmt"


// No se Exporta porque empieza con minúscula
func print_variables() {
	var a string = "Hello"
	var b int = 42
	var c float64 = 3.14
	var d bool = true

	fmt.Println("String:", a)
	fmt.Println("Integer:", b)
	fmt.Println("Float:", c)
	fmt.Println("Boolean:", d)
}

// Se Exporta porque empieza con mayúscula
func PrintVariables() {
	var a string = "Hello"
	var b int = 42
	var c float64 = 3.14
	var d bool = true

	fmt.Println("String:", a)
	fmt.Println("Integer:", b)
	fmt.Println("Float:", c)
	fmt.Println("Boolean:", d)
}