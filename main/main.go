package main

import (
	"fmt"

	"github.com/par97/goutil/color"
)

func main() {
	color.PrintTerminalColors()
	color.PrintTerminalStyles()
	color.PrintTerminalBackgroundColors()

	fmt.Println("")
	color.Println("hello", color.Green, color.Italic)
	color.Println("world", color.Red, color.Bold)
	color.Println("wow!", color.Blue, color.Underline)

	fmt.Println("")
	color.Print("hello ", color.Green, color.Italic)
	color.Print("world ", color.Red, color.Bold)
	color.Println("wow!", color.Blue, color.Underline)
}
