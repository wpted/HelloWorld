package hello

import (
	"strings"
)

const (
	EnglishCommonPrefix = "Hello, "
	ChineseCommonPrefix = "你好, "
	FrenchCommonPrefix  = "Bonjour, "
)

// greetingPrefix is a internal helper function for Hello()
func greetingPrefix(language string) (prefix string) {
	switch language {
	case "english":
		{
			prefix = "Hello, "
		}
	case "traditional chinese":
		{
			prefix = "你好, "
		}
	case "french":
		{
			prefix = "Bonjour, "
		}
	case "japanese":
		{
			prefix = "こんにちは, "
		}
	default:
		{
			prefix = "Hello, "
		}
	}

	return
}

// Hello says hello to the input user, defaulted to the world in different languages
func Hello(name, language string) string {
	greetings := ""
	language = strings.ToLower(language)
	if name == "" {
		switch language {
		case "english":
			{
				greetings = "Hello, World"
			}
		case "traditional chinese":
			{
				greetings = "你好, 世界"
			}
		case "french":
			{
				greetings = "Bonjour, Monde"
			}
		case "japanese":
			{
				greetings = "こんにちは, 世界"
			}
		default:
			{
				greetings = "Hello, World"
			}
		}
	} else {
		if language == "japanese" {
			name += "さん"
		}
		greetings = greetingPrefix(language) + name
	}

	return greetings

}
