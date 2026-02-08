package gentleman

import "fmt"

func PrintTypes() {

	// Arrays y Slices
	arrayFijo := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array fijo:", arrayFijo)
	sliceVariable := []int{1, 2, 3, 4, 5}
	sliceVariable = append(sliceVariable, 6)
	fmt.Println("Slice variable:", sliceVariable)

	// Mapas
	diccionario := map[string]int{
		"El alquimista": 1,
		"El principito": 2,
	}
	fmt.Println("Diccionario:", diccionario)
	

	// Structs
	type Persona struct {
		Nombre string
		Apellido string
	}

	p1 := Persona{
		Nombre: "Cristopher",
		Apellido: "Angulo",
	}
	fmt.Println("Persona:", p1)


}
