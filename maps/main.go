package main

import "fmt"

func main() {
	// colors := make(map[string]string)

	// var colors map[string]string

	colors := map[string]string{
		"red":  "#ff0000",
		"blue": "#00ff00",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for", value, "is", key)
	}
}
