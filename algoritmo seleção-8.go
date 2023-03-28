package main

import "fmt"

func main() {
	var x int

	fmt.Print("Digite um número inteiro entre 1 e 7: ")
	fmt.Scanln(&x)

	if x == 1 {
		fmt.Println("Domingo")
	} else if x == 2 {
		fmt.Println("Segunda-feira")
	} else if x == 3 {
		fmt.Println("Terça-feira")
	} else if x == 4 {
		fmt.Println("Quarta-feira")
	} else if x == 5 {
		fmt.Println("Quinta-feira")
	} else if x == 6 {
		fmt.Println("Sexta-feira")
	} else if x == 7 {
		fmt.Println("Sábado")
	} else {
		fmt.Println("Número inválido")
	}
}
