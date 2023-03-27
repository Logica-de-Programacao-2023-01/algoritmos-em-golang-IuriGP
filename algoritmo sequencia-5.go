package main

import "fmt"

func main() {
	var idade int
	fmt.Print("Qual a sua idade? ")
	fmt.Scan(&idade)
	fmt.Print("Sua idade em dias é: ", idade*365)
}
