package strings

import "fmt"

// Wrap returns the provided string wrapped in the provided wrapper
func Wrap(str string, wrapper string) string {
	wrappedString := fmt.Sprintf("%v%v%s", wrapper, str, wrapper)
	return wrappedString
}
