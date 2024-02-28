package main

import "fmt"

func main() {
	myBill := newBill("mario's bill")

	fmt.Println(myBill)
	myBill.updateTip(10)
	myBill.addItem("onion soup", 4.50)
	myBill.addItem("veg pie", 8.95)
	myBill.addItem("toffee pudding", 4.95)
	myBill.addItem("coffee", 3.25)

	fmt.Println(myBill.format())
}
