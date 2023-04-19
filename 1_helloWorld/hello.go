package hello

const (
	englishLang = "English"
	spanishLang = "Spanish"
	russianLang = "Russian"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	russianHelloPrefix = "Привет, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return getPrefix(language) + name
}

// (prefix string) - creates a variable with name + variable name in return statement could be omitted
// Also it's self documeting
func getPrefix(language string) (prefix string) {
	switch language {
	case spanishLang:
		prefix = spanishHelloPrefix
	case russianLang:
		prefix = russianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

// TDD concepts:
// 1. Write a failing test and see it fail so we know we have written a relevant test for our requirements and seen that it produces an easy to understand description of the failure
// 2. Writing the smallest amount of code to make it pass so we know we have working software
// 3. Then refactor, backed with the safety of our tests to ensure we have well-crafted code that is easy to work with