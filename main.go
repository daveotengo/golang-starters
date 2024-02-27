package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	phoneBook := map[int]string{
		123213123: "David",
		243243242: "Kwadwo",
		546456454: "Emmanuel",
		122324233: "Reginald",
	}

	fmt.Println(phoneBook)
	phoneBook[243243242] = "Abena"
	fmt.Println(phoneBook)

}
