package main

import "fmt"

type englishBot struct{}
type frenchBot struct{}

func main() {
	eb := englishBot{}
	fb := frenchBot{}

	printGreeting(eb)
	printGreeting(fb)
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(fb frenchBot) {
	fmt.Println(fb.getGreeting())
}

func (eb englishBot) getGreeting() string {
	//custom logic
	return "hi"
}

func (fb frenchBot) getGreeting() string {
	//custom logic
	return "bonjour"
}
