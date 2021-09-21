package main

import "fmt"

var area int

func main() {
	var a int
	var b int

	fmt.Print("Enter the length :")
	fmt.Scan(&a)
	fmt.Print("Enter the width :")
	fmt.Scan(&b)
	area = a * b

	if a != b {
		fmt.Print("it is not a square")

	} else {

		fmt.Print("area of Square :", area)

	}

}
