package cmd

import "fmt"

const s string = "Hello, World!"
const n = 10000000

func GetString() string {
	return s
}

func PrintConstants() {
	fmt.Println("Constant String:", s)
	fmt.Println("Constant Number:", n)
}