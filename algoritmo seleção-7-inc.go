package main

import "fmt"

func main() {

	var x float64
	fmt.Print("Qual valor do salário?")
	fmt.Scan(&x)
	if x >= 1000 {
		print("O funcionário recebera um aumento de 5%: ", 0.05*x)
	} else if x < 1000 {
		print("O funcionário recebera um aumento de 10%: ", 0.10*x)
	}
}
