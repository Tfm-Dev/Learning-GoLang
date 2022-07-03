package main

import "fmt"

func main() {
	//Printf with any value
	fmt.Printf("Any value: %v\n", "Teste One")
	fmt.Printf("Any value: Teste %v\n", 2)
	fmt.Printf("Any value: Teste %v\n", 2.1)
	fmt.Printf("Any value: Teste Three %v\n\n", true)

	//PrintF with type value
	fmt.Printf("Type value: %T\n", "Teste")
	fmt.Printf("Type value: %T\n", 1)
	fmt.Printf("Type value: %T\n", 1.1)
	fmt.Printf("Type value: %T\n\n", true)

	//PrintF literal %
	fmt.Printf("%%\n\n")

	//PrintF Data Type
	fmt.Printf("String %s or %q\n", "Thiago", "Moura")
	fmt.Printf("Integer %b or %o or %d or %x\n", 2, 8, 999999999999999999, 999999999999999999)
	fmt.Printf("Float %f or %.2f\n", 1.1, 1.1241)
	fmt.Printf("Boolean %t\n", true)
}
