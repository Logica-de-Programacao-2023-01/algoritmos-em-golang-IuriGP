package main

import "fmt"

func main() {
	var x int
	fmt.Println("Qual valor de x?")
	fmt.Scan(&x)
	fmt.Print(
		"O dobro desse número é: ", x*2,
		"; O triplo desse número é: ", x*3,
		"; O quaduplo desse número é: ", x*4)
}
