// Package main is a simple hello world application
package main

import (
	"fmt"

	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
)

// GenText generates the text "Hello World!"
func GenText() string {
	return "Hello World!"
}

func main() {
	fmt.Println(GenText())
	say, _ := cowsay.Say(
		"Muuh!",
		cowsay.Type("default"),
		cowsay.BallonWidth(40),
	)
	fmt.Println(say)
}
