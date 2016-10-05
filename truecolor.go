package main

import "fmt"

var raw = `/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/`

func main() {
	for c := 0; c < 77; c++ {
		r := 255 - (c * 255 / 76)
		g := (c * 510 / 76)
		b := (c * 255 / 76)
		if g > 255 {
			g = 510 - g
		}
		fmt.Printf("\033[48;2;%d;%d;%dm", r, g, b)
		fmt.Printf("\033[38;2;%d;%d;%dm", 255-r, 255-g, 255-b)
		fmt.Printf("%s\033[0m", substr(raw, c+1))
	}
	fmt.Println()
}

func substr(s string, c int) string {
	if c >= len(s) {
		return s[c:len(s)]
	}
	if c < len(s) {
		return s[c : c+1]
	}
	panic("should not happen")
}
