package main

import "fmt"

func main() {
	//Define var
	var name string
	var idade uint

	//Set values
	name = "Thiago"
	idade = 21

	//Define and set values
	altura := 1.83
	dev := true

	fmt.Println(name, idade, "Anos")
	fmt.Println(altura)
	fmt.Println("Desenvolvedor? ", dev)
}
