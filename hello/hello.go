package main

import "fmt"

const (
	spanish                   = "Spanish"
	spanishHelloPrefix string = "Hola, "
	french                    = "French"
	frenchHelloPrefix  string = "Bonjour, "

	englishHelloPrefix string = "Hello, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	}
	if language == french {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
