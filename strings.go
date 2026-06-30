package strings

import "fmt"

// Wrap returns the provided string wrapped in the provided wrapper
func Wrap(str string, wrapper string) string {
	wrappedString := fmt.Sprintf("%v%v%s", wrapper, str, wrapper)
	return wrappedString
}

func Prefix(str string, prefix string) string {
	prefixedString := fmt.Sprintf("%v%v", prefix, str)
	return prefixedString
}

func Suffix(str string, suffix string) string {
	suffixedString := fmt.Sprintf("%v%v", str, suffix)
	return suffixedString
}
