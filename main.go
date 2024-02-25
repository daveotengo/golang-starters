package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greetings := "Hello there friends"
	fmt.Println(strings.Contains(greetings, "Hello"))
	fmt.Println(strings.ReplaceAll(greetings, "Hello", "Hi"))
	fmt.Println(strings.Index(greetings, "th"))
	fmt.Println(strings.Split(greetings, " "))

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	sort.Strings(names)
	fmt.Println(names)

	index2 := sort.SearchStrings(names, "bowser")
	fmt.Println(index2)

}
