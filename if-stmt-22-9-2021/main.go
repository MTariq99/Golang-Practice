package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Print("Enter the first value :")
	fmt.Scan(&a)
	fmt.Print("Enter the second value :")
	fmt.Scan(&b)

	if a > b {
		fmt.Print(a, " > ", b)
	} else {
		fmt.Print(b, " > ", a)
	}
}
