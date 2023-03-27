package main

import "fmt"

func main() {

	var x float64
	var libras = 0.45

	fmt.Print("Quantos quilose você pesa?")
	fmt.Scan(&x)
	fmt.Print("Esse é o valor em libras de seu peso: ", x/libras)
}
