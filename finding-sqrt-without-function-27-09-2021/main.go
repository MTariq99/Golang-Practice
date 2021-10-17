package main

import "fmt"

func main() {
	var a int
	fmt.Print("enter the number")
	fmt.Scan(&a)

	if a < 0 {
		fmt.Print("Error: Negative Number!")

	}

}
