package main

import "fmt"

func main() {
	var ages = []int{23, 45, 34}

	var names = [4]string{"Ama", "Kofi", "Kwadwo", "Afua"}

	ages[2] = 2
	fmt.Println(ages, len(ages))

	ages = append(ages, 86)

	fmt.Println(ages, len(ages))

	fmt.Println(names, len(names))

	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

}
