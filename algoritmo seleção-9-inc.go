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
	if x > y && y > z {
		fmt.Println(x, y, z)
	} else if x > y && y < z {
		fmt.Println(x, z, y)
	} else if y > x && x > z {
		fmt.Println(y, x, z)
	} else if y > x && x < z {
		fmt.Println(y, z, x)
	} else if z > y && y > x {
		fmt.Println(z, y, x)
	} else if z > x && y < x {
		fmt.Println(z, x, y)
	}
}
