package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

// all types that implement a getGreeting function that returns a type string
// are now also honorary members of type bot
type bot interface {
	getGreeting()
	// eg complex example with multiple args and multiple return types
	// getGreeting(string, int) (string, error)
}

// can omitt the value of the receiver as it's not used hence not (eb englishBot) getGreeting...
func (englishBot) getGreeting() string {
	// some examlpe custom ENGLISH BOT code
	return "Hello World"
}

func (spanishBot) getGreeting() string {
	// some examlpe custom Spanish BOT code
	return "Hola!"
}

//  Here be dragons
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }
// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
