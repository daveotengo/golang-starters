package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("value of i is:", i)
	}

	names := []string{"mario", "luigi", "yooshi", "peach"}

	for i := 0; i < len(names); i++ {
		fmt.Println("the value of i is:", names[i])
	}

	for index, value := range names {
		fmt.Printf("The value of index %v is  %v", value, index)
	}

	for _, value := range names {
		fmt.Printf("The value is  %v", value)
	}

}
