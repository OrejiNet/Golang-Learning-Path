package cmd

import "fmt"

func Arrays() {

	var a [5]int
	fmt.Println("Array a:", a)

	a[4] = 100
	fmt.Println("Array a after assignment:", a)
	fmt.Println("Length of array a:", len(a))
	fmt.Println("Element at index 4:", a[4])

	// Declare and initialize an array of integers
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr)

	b := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array b:", b)

	var twoD [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("2D Array:", twoD)

	var twoD2 [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD2[i][j] = i + j
		}
	}
	fmt.Println("2D Array twoD2:", twoD2)
}
