package main

import "fmt"

func main() {
	// map is without orders
	m := map[string]string{
		"name":   "jack",
		"course": "go",
	}
	m1 := make(map[string]int)
	fmt.Println(m, m1)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting value")
	name, ok := m["name"]
	fmt.Println(name, ok)
	course, ok := m["coourse"]
	fmt.Println(course, ok)

	fmt.Println("Deleting value")
	name, ok = m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
