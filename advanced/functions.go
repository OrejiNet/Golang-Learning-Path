package advanced


func Sum(a, b int) int {
	return a + b
}

func PlusPlus(a, b, c int) int {
	return a + b + c
}

func Values() (int, int) {
	return 3, 4
}

func Variadicas(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num			
	}
	return sum
}

