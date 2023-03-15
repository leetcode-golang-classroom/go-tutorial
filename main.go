package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(promt string, r *bufio.Reader) (string, error) {
	fmt.Print(promt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill name:", reader)
	b := newBill(name)
	fmt.Println("Created the bill -", b.name)
	return b
}
func promtOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose a option (a - add item, s - save bill, t - add tip):", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		fmt.Println(name, price)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		fmt.Println(tip)
	case "s":
		fmt.Println("you choose s")
	default:
		fmt.Println("that was not valid option...")
		promtOptions(b)
	}

}
func main() {
	myBill := createBill()
	promtOptions(myBill)
}
