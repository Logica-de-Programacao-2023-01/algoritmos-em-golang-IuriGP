package main

import "fmt"

func main() {
	var salario float32
	fmt.Print("Qual o salário do fnucionário? ")
	fmt.Scan(&salario)
	var aumento = salario * 0.15
	fmt.Print("O aumento de salário desse funcionario será de: ", salario+aumento)
}
