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

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
