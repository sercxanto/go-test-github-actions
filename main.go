package main

import (
	"fmt"

	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
)

func genText() string {
	return "Hello World!"
}

func main() {
	fmt.Println(genText())
	say, _ := cowsay.Say(
		"Hello World!",
		cowsay.Type("default"),
		cowsay.BallonWidth(40),
	)
	fmt.Println(say)
}
