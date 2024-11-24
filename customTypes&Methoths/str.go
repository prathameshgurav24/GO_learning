package main

import "fmt"

type str string //str is custom type made out of string

func (text str) log() {
	fmt.Print(text)
}

func main() {
	var name str = "Virat"
	name.log()
}
