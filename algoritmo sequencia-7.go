package main

import "fmt"

func main() {
	var x float32
	fmt.Print("Qual valor numérico? ")
	fmt.Scan(&x)
	x++
	fmt.Print("O sucessor desse número é: ", x)
	fmt.Print(" E o antecessor desse número é: ", x-2)
}
