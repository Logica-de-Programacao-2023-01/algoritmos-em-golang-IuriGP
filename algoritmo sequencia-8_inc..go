package main

import (
	"fmt"
)

func main() {
	var dias float32
	var diario float32
	fmt.Print("Quantos dias o funcionario trabalhou? ")
	fmt.Scan(&dias)
	fmt.Print("Quanto esse funiconario ganhar por dia? ")
	fmt.Scan(&diario)
	var horas = dias * 24
	fmt.Print("o valor do salário recebido pelo funcionario é de: ", horas*diario)
}
