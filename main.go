package main

import "fmt"

func main() {
	var ageOne = 12

	var nameOne = "Ama"

	fmt.Printf("my age is %v and my name is %v \n", ageOne, nameOne)
	fmt.Printf("my age is %q and my name is %q \n", ageOne, nameOne)
	fmt.Printf("age is of type %T \n", ageOne)
	fmt.Printf("You scored %f points \n", 225.55)
	fmt.Printf("You scored %0.1f points \n", 225.55)

	var str = fmt.Sprintf("my age is %v and my name is %v", ageOne, nameOne)

	fmt.Println("the saved string is", str)
}
