package main

import (
	"fmt"
)

const (
	french     = "french"
	portuguese = "portuguese"
	spanish    = "spanish"

	englishHelloPrefix    = "hello, "
	frenchHelloPrefix     = "bonjour, "
	portugueseHelloPrefix = "olá, "
	spanishHelloPrefix    = "hola, "
)

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
