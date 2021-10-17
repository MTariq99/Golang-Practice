package main

import "fmt"

func main() {

	for i := 0; i <= 7; i++ {
		for j := 0; j <= 25; j++ {
			if i%2 == 0 && i <= 5 {
				if j <= 7 {
					fmt.Printf("* ")
				} else {
					fmt.Printf("=")
				}
			} else if i%2 != 0 && i <= 5 {
				if j <= 7 {
					fmt.Printf(" *")
				} else {
					fmt.Printf("=")
				}

			} else if i > 5 {

				fmt.Print("=")
			}

		}
		fmt.Println(" ")
	}

}
