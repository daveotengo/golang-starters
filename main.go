package main

import "fmt"

func updateName(x string) string {
	x = "wedge"
	return x
}

func main() {
	name := "tifa"
	name = updateName(name)

	fmt.Println(name)
}
