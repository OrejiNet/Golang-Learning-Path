package cmd

import "fmt"

func ForLoop() {

	j := 1
	//var j int = 1
	for j < 10 {
		fmt.Println("Value of j:", j)
		j++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	for {
		fmt.Println("This will run indefinitely")
		break // Agrega un break para evitar un bucle infinit
	}

	for n := 0; n < 10; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println("Odd number:", n)
	}
}