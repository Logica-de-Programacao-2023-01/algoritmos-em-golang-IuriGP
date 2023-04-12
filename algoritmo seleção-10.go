package main

import "fmt"

func main() {
	var x float32
	fmt.Print("Qual a sua idade: ")
	fmt.Scan(&x)

	fmt.Println("Sua classificação etária é: ", x)
	if x >= 18 {
		fmt.Println(" Adulto ")
	} else if x >= 14 && x <= 17 {
		fmt.Println(" juvenil ")
	} else if x > 10 && x <= 13 {
		fmt.Println(" infantil ")
	} else {
		fmt.Println(" mirin ")
	}
}
