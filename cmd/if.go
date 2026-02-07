package cmd

import "fmt"

func IfStatements() {
	x := 10
	y := 20

	if x < y {
		fmt.Println("x is less than y")
	} else if x > y {
		fmt.Println("x is greater than y")
	} else {
		fmt.Println("x is equal to y")
	}

	if num := 8; num%2 == 0 {
		fmt.Println(num, "is even")
	}
}