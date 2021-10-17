package main

import "fmt"

func main() {

	var a int = 5
	var b int
	for i := 1; i <= 5; i++ {
		b = 0
		for j := 1; j <= a-i; j++ {
			fmt.Print(" ")
		}
		for {
			fmt.Print("*")
			b++
			if b == 2*i-1 {
				break
			}
		}
		fmt.Println("")
	}

}
