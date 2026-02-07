package cmd

import (
	"fmt"
	"time"
)

func SwitchStatements() {
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("It's the start of the week!")
	case "Tuesday":
		fmt.Println("It's the middle of the week!")
	case "Wednesday":
		fmt.Println("It's the middle of the week!")
	case "Thursday":
		fmt.Println("It's almost the end of the week!")
	case "Friday":
		fmt.Println("It's the end of the week!")
	case "Saturday":
		fmt.Println("It's the end of the week!")
	case "Sunday":
		fmt.Println("It's the end of the week!")
	default:
		fmt.Println("It's the end of the week!")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	t := time.Now()
	fmt.Println(t)
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 18:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	whatamI := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Println("I'm an int")
		case string:
			fmt.Println("I'm a string")
		default:
			fmt.Printf("Don't know type %T\n", v)
		}
	}

	whatamI(42)
	whatamI("Hello")
	whatamI(3.14)
}