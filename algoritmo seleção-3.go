package main

import "fmt"

func main() {
	var x int
	fmt.Println("qual número?")
	fmt.Scan(&x)

	if x%2 == 0 {
		fmt.Println("é um número par")

	} else {
		fmt.Println("é um número ímpar")
	}

}
