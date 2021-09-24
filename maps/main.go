package main

import "fmt"

func main() {
	colors := make(map[string]string)

	// var colors map[string]string

	// colors := map[string]string{
	// 	"red":  "#ff0000",
	// 	"blue": "#00ff00",
	// }

	colors["white"] = "#ffffff"

	colors["black"] = "#000000"

	fmt.Println(colors)

	delete(colors, "white")

	fmt.Println(colors)
}
