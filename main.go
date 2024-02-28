package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Create a new bill name: ")
	name, _ := getInput("Create New bill name:", reader)
	// name = strings.TrimSpace(name)
	b := newBill(name)
	fmt.Println("Created new bill -", b.name)
	return b
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
	//fmt.Println(myBill)
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a -add item s -save bill, t - add tip): ", reader)
	fmt.Println(opt)

	switch opt {
	case "a":
		name, _ := getInput("Item name", reader)
		price, _ := getInput("Item price", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("item added", name, price)
	case "t":
		tip, _ := getInput("Enter tip amount ($):", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println(tip)
		promptOptions(b)

	case "s":
		b.save()
		fmt.Println("You saved file -", b.name)
	default:
		fmt.Println("that was not a valid option")
		promptOptions(b)
	}
}

func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("bill was saved to file")
}
