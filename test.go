package main

import (
	"fmt"
)

func main() {
	var x float32
	var y float32
	fmt.Print("Quantos dias o funcionario trabalhou? ")
	fmt.Scan(&x)
	fmt.Print("Quanto esse funiconario ganhar por dia? ")
	fmt.Scan(&y)
	var mes = x / 30
	fmt.Print("o valor do salário recebido pelo funcionario é de: ")
}
