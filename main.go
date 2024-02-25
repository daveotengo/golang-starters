package main

import "fmt"

func main() {
	age := 45

	if age < 30 {
		fmt.Println("age is less than 30 ")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is equal 45")
	}

	names := []string{"kwasi", "kojo", "abena", "yaw", "edward"}
	for index, value := range names {
		if index == 1 {
			fmt.Println("continue at pos ", index)
			continue
		}
		fmt.Printf("the value at pos %v is %v \n", index, value)
	}

}
