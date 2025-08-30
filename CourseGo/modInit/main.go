package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello from local module")
	fmt.Println(quote.Go())
}
