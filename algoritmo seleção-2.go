package main

import "fmt"

func main() {
	x := 50
	y := 10
	z := 15
	if x > y && y > z {
		fmt.Println("z é o menor número")
	} else if x > y && y < z {
		fmt.Println("y é o menor número")
	} else {
		fmt.Println("x é o menor número")
	}
}
