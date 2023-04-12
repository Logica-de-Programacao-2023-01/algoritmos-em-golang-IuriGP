package main

import (
	"fmt"
)

func main() {
	var x = 1
	for i := 1; i <= 100; i++ {
		fmt.Print(x * i)
		if i%5 == 0 && i%3 == 0 {
			fmt.Println(" FizzBuzz ")
		} else if i%5 == 0 {
			fmt.Println(" buzz ")
		} else if i%3 == 0 {
			fmt.Println(" Fizz ")
		} else {
			fmt.Println(" ")
		}
	}
}
