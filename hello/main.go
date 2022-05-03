package main

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const spanish = "Spanish"
const frenchHelloPrefix = "Bonjour, "
const french = "French"

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := greetingPrefix(language)

	return prefix + name
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

}
