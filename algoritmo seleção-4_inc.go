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
	var peso_final = x * peso_2
	fmt.Println("Seu peso é: ", peso_final)
	if peso_final < 18 && peso_final < 25 {
		fmt.Println("Você está abaixo da média")
	} else if peso_final > 18 && peso_final < 25 {
		fmt.Println("Você está na média")
	} else {
		fmt.Println("Você está acima da média")
	}
}
