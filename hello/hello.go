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

	prefix := englishHelloPrefix
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
