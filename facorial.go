package main

import "fmt"

func main() {
	factorial := 1
	for i := 4; i > 0; i-- {
		factorial = factorial * i
	}
	fmt.Println(factorial)

}
