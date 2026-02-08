package main

import (
	"curso-go/advanced"
	"curso-go/cmd"
	"curso-go/gentleman"
	"fmt"
	"sync"
	"time"
)

func main() {
	//main_concepts()
	//advanced_functions()
	//level3()
	//memoria()
	//interfaces()
	channels()
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


func channels() {
	canal := make(chan string)
	go gentleman.CallHello(canal)
	gentleman.PrintMessage(canal)

	canal2 := make (chan int)
	go func(){
		for i := range 5 {
			canal2 <- i
		}
		close(canal2)
	}()

	for num := range canal2 {
		fmt.Println("Numero recibido:", num)
	}

	// Mutex Bloquear y desbloquear el acceso a un recurso compartido para evitar condiciones de carrera.
	var contador int
	var mutex sync.RWMutex

	// Writer - Incrementa el contador
	go func(){
		for i := 0; i < 5; i++ {
			mutex.Lock()
			contador++
			mutex.Unlock()
			time.Sleep(100 * time.Millisecond)
		}
	}()
	
	// Reader - Lee el valor del contador
	for i := 0; i < 5; i++ {
		go func(id int ){
			for j := 0; j < 5; j++ {
				mutex.RLock()
				fmt.Printf("Leyendo desde Goroutine %d - Contador: %d\n", id, contador)
				mutex.RUnlock()
				time.Sleep(150 * time.Millisecond)
			}
		}(i)
	}

	time.Sleep(2 * time.Second)

}