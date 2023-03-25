package main

import "fmt"

func main() {
	var x int
	var y int
	var z int
	fmt.Println("qual valor de x?")
	fmt.Scan(&x)
	fmt.Println("qual valor de y?")
	fmt.Scan(&y)
	fmt.Println("Qual valor de z?")
	fmt.Scan(&z)
	fmt.Println("A soma de x+y+z é=", x+y+z)

}
