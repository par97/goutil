package color

import "fmt"

type Color uint8

const Black Color = 30
const Red Color = 31
const Green Color = 32
const Yellow Color = 33
const Blue Color = 34
const Magenta Color = 35
const Cyan Color = 36
const White Color = 37

type Style uint8

const Normal Style = 0
const Bold Style = 1
const Faint Style = 2
const Italic Style = 3
const Underline Style = 4
const SlowBlink Style = 5
const RapidBlink Style = 6

const RESET = "\033[0m"

func PrintTerminalColors() {
	fmt.Print("font colors: ")
	for i := 30; i < 40; i++ {
		fmt.Printf("\033[%dm%d%s ", i, i, RESET)
	}
	fmt.Println("")
}

func PrintTerminalStyles() {
	fmt.Print("font styles: ")
	for i := 0; i < 8; i++ {
		fmt.Printf("\033[39;%dm%2d%s ", i, i, RESET)
	}
	fmt.Println("")
}

func PrintTerminalBackgroundColors() {
	fmt.Print("bkgd colors: ")
	for i := 40; i < 50; i++ {
		fmt.Printf("\033[%dm%d%s ", i, i, RESET)
	}
	fmt.Println("")
}

func Println(msg string, color Color, style Style) {
	fmt.Printf("\033[%d;%dm%s%s\n", color, style, msg, RESET)
}

func Print(msg string, color Color, style Style) {
	fmt.Printf("\033[%d;%dm%s%s", color, style, msg, RESET)
}
