package main

import "fmt"

func main() {
	// var colors map[string]string

	// colors := make(map[string]string)
	// colors["red"] = "#ff0000"
	// colors["green"] = "#00ff00"

	colors := map[string]string{
		"black": "#000000",
		"white": "#ffffff",
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Hex code for %s is %s\n", color, hex)
	}
}
