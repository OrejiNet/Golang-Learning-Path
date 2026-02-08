package main

import (
	"curso-go/cmd"
	"fmt"
	"curso-go/advanced"
	"curso-go/gentleman"
)

func main() {
	//main_concepts()
	//advanced_functions()
	//level3()
	//memoria()
	interfaces()
}

func main_concepts() {
	fmt.Println("Go" + "Lang")
	fmt.Println("1+1=", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(!true)
	fmt.Print(true || false, "\n")
	cmd.PrintVariables()
	cmd.PrintConstants()
	cmd.ForLoop()
	cmd.IfStatements()
	cmd.SwitchStatements()
	cmd.Arrays()
	cmd.ShowMap()
}


func advanced_functions(){
	fmt.Println("Sum of 1 and 2 is:", advanced.Sum(1,2))
	fmt.Println("Sum of 1, 2 and 3 is:", advanced.PlusPlus(1,2,3))
	advancedValues1, advancedValues2 := advanced.Values()
	fmt.Println("Values returned by Values() function:", advancedValues1, advancedValues2)
	fmt.Println("Sum of 1, 2, 3 and 4 is:", advanced.Variadicas(1,2,3,4))
	nums:= []int{1,2,3,4,5}
	fmt.Println("Sum of nums slice is:", advanced.Variadicas(nums...))	
	seq := advanced.IntSequence()
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println("Factorial of 5 is:", advanced.Factorial(5))
}


func level3() {
	result, err := gentleman.Dividir(10, 0)
	
	if err != nil { fmt.Println("Error:", err)
	} else { fmt.Println("Resultado:", result)}
	
	gentleman.PrintNombres("Cristopher", "Angulo", "Gonzalez")
	gentleman.CallContador(5)

	rectangulo := gentleman.Rectangulo{Ancho: 5, Alto: 3}
	gentleman.InfoRectangulo(rectangulo)

	rectangulo2 := gentleman.Rectangulo{Ancho: 10, Alto: 4}
	rectangulo2.Area()
	rectangulo2.Perimetro()
}

func memoria(){
	gentleman.TestPunteros()
}

func interfaces(){
	gentleman.CallInfo()
}