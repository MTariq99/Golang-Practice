package main

import "fmt"

func main() {
	var a int
	var b int
	var c int
	var d int
	d = 2 ^ 2
	X := -b + (b*b-4*a*c)*1/2
	Y := -b - (b*b-4*a*c)*1/2

	C := X + Y

	fmt.Print("Enter the value of a :", d)
	fmt.Scan(&a)
	fmt.Print("Enter the value of b :")
	fmt.Scan(&b)
	fmt.Print("Enter the value of c :")
	fmt.Scan(&c)

	fmt.Println(C)

}
