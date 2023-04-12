package main

import "fmt"

func main() {
	var x int
	fmt.Print("Qual valor de x?")
	fmt.Scan(&x)
	for i := 0; i <= 10; i++ {
		fmt.Println("O resultado é: ", x*i)
	}
}
