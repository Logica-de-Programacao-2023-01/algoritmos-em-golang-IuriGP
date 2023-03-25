package main

import "fmt"

func main() {

	var x float32
	var y float32

	fmt.Print("Qual o seu peso(em KG)? ")
	fmt.Scan(&x)
	fmt.Print("Qual a sua Altura?")
	fmt.Scan(&y)
	var peso_2 = y * y
	fmt.Println("Seu peso é: ", x/peso_2)
}
