package main

import "fmt"

func main() {

	var x int
	var y int
	fmt.Print("Qual valor de x?")
	fmt.Scan(&x)
	fmt.Print("qual valor de y?")
	fmt.Scan(&y)
	if x > 0 && y > 0 {
		fmt.Print("A multiplicação desses dois valores é: ", x*y)
	} else if x < 0 || y < 0 {
		fmt.Print("A soma desses dois valores é: ", x+y)
	}
}
