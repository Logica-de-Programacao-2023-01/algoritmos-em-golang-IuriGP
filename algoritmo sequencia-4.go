package main

import "fmt"

func main() {

	var x int
	var y int
	var z int

	fmt.Print("Qual valor de x?")
	fmt.Scan(&x)
	fmt.Print("qual valor de y?")
	fmt.Scan(&y)
	fmt.Print("Qual valor de z?")
	fmt.Scan(&z)
	fmt.Print("A média ponderada de Peso_2 é de: ", (x+y+z)*2/3,
		" ;A média ponderada de Peso_3 é de: ", (x+y+z)*3/3,
		" ;A média ponderada de Peso_5 é de: ", (x+y+z)*5/3)
}
