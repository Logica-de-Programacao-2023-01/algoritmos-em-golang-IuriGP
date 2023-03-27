package main

import "fmt"

func main() {

	var x float32
	var y float32

	fmt.Print("Qual o seu peso(em KG)? ")
	fmt.Scan(&x)
	fmt.Print("Qual a sua Altura?")
	fmt.Scan(&y)
	var peso = y * y
	fmt.Println("Seu IMC é: ", x*peso)
}
