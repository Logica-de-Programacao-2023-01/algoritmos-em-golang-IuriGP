package main

import "fmt"

func main() {
	var x int
	fmt.Print("Qual valor de x?")
	fmt.Scan(&x)
	if x%3 == 0 && x%5 == 0 {
		fmt.Println("Esse número é divisivel por 3 e 5")
	} else if x%3 != 0 && x%5 == 0 {
		fmt.Println("Esse número é divisivel só 5")
	} else if x%3 == 0 && x%5 != 0 {
		fmt.Println("Esse número é divisivel só por 3")
	} else {
		fmt.Println("Esse numero não é divisivel por 3 nem por 5")
	}

}
