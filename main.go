package main

import (
	"fmt"

	"github.com/saish24/greetings"
)

func main() {
	if str, err := greetings.RandomQuote(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(str)
	}
}
