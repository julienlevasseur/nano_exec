package main

import (
	"fmt"
)

var input string
var green string = "\033[32m"
var yellow string = "\033[93m"
var red string = "\033[91m"
var orange string = "\033[38;5;208m"
var end string = "\033[0m"

func PrintGreen(input string) {
	fmt.Printf("%s%s%s", green, input, end)
}

func PrintYellow(input string) {
	fmt.Printf("%s%s%s", yellow, input, end)
}

func PrintRed(input string) {
	fmt.Printf("%s%s%s", red, input, end)
}

func PrintOrange(input string) {
	fmt.Printf("%s%s%s", orange, input, end)
}
