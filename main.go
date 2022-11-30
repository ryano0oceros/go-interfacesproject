package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type frenchBot struct{}

func main() {
	eb := englishBot{}
	fb := frenchBot{}

	printGreeting(eb)
	printGreeting(fb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	//custom logic
	return "hi"
}

func (fb frenchBot) getGreeting() string {
	//custom logic
	return "bonjour"
}

//no changes through 57...
