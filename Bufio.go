package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Define scanner
	scanner := bufio.NewScanner(os.Stdin)

	//Get input
	scanner.Scan()

	//Parse text to int
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	//Print value
	fmt.Printf("%d", input)
}
