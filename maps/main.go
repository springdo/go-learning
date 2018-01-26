package main

import "fmt"

func main() {
	// 1,	map with keys of type string and values of type string
	// variableName := map[string]string{"white": "ffffff"}

	colors := map[string]string{
		"red":   "#ff0000",
		"white": "#ffffff",
		"green": "GREEN_HEX",
	}
	fmt.Println(colors)

	// 2, create an empty map
	var colors2 map[string]string
	// or another way for empty map and then assign key / vals to it
	colors22 := make(map[string]string)
	colors22["red"] = "RED_HEX"

	fmt.Println(colors2)
	fmt.Println(colors22)

	// to delete items in a map
	delete(colors, "white")
	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for key, val := range c {
		fmt.Printf(" %s : %s \n", key, val)
	}
}
