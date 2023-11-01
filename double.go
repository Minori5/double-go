package main

import (
	"fmt"
)

func double(num *int) {
	*num *= 2
}

func main() {
	amount := 10
	double(&amount)
	fmt.Println(amount)
}
