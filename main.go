package main

import (
	"bufio"
	"fmt"
	"os"
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
}
