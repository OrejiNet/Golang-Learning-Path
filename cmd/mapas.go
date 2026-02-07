package cmd

import "fmt"

func ShowMap() {

	m := make(map[string]int)
	m["Alice"] = 30
	m["Bob"] = 25
	m["Charlie"] = 35
	fmt.Println("Map m:", m)

	delete(m, "Bob")
	fmt.Println("Map m after deleting Bob:", m)

	_, prs:= m["Bob"]
	fmt.Println("Is Bob in the map?", prs)

	n := map[string]int{
		"Dave": 40,
		"Eve":  28,
	}
	fmt.Println("Map n:", n)
}