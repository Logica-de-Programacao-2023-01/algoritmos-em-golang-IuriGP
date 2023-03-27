package main

import "fmt"

func main() {

	var x float32
	var y float32

	fmt.Print("Qual preço do produto?")
	fmt.Scan(&x)
	fmt.Print("qual valor do descontado?")
	fmt.Scan(&y)
	var desconto = y / 100
	var valor_descontado = x * desconto
	fmt.Print("O valor descontedo é de: ", x-valor_descontado)
}
