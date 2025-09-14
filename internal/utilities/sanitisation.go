package utilities

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// input is the sting you want to sanitise, fieldName is used if there was a banned character
//
// e.g backslash n is banned, if there was one in a title string this function would
// print the following to the console
//
// Banned characters were detected in the title. They have been removed.
// To see the full list of banned characters, please view the documentation at github.com/samhodg1993/toto-todo-cli
func SanitizeInput(input string, fieldName string) string {
	// Remove ANSI escape sequences including \033[C and other variants
	// This covers ESC[ followed by any combination of digits, semicolons, and letters
	ansiRegex := regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)
	cleaned := ansiRegex.ReplaceAllString(input, "")

	// Also handle alternative escape sequence formats
	// \033[C, \e[C, ^[[C etc.
	// This is essential as the flag -C is used to clear the terminal,
	// not removing this would mean that if a title was do something
	// in the same way that the toto ls -C command works,
	// then the screen would clear after printing
	altEscapeRegex := regexp.MustCompile(`\\0?33\[[0-9;]*[a-zA-Z]|\\e\[[0-9;]*[a-zA-Z]|\^\[\[[0-9;]*[a-zA-Z]`)
	cleaned = altEscapeRegex.ReplaceAllString(cleaned, "")

	// Remove other control characters (ASCII 0-31 except tab and newline if needed)
	// Keep only printable characters and common whitespace
	var result strings.Builder
	for _, r := range cleaned {
		if unicode.IsPrint(r) || r == ' ' || r == '\t' || r == '\n' {
			result.WriteRune(r)
		}
	}

	sanitisedInput := strings.TrimSpace(result.String())

	if sanitisedInput != input {
		fmt.Printf("Banned characters were detected in the %s. They have been removed.", fieldName)
		fmt.Println("To see the full list of banned characters, please view the documentation at github.com/samhodg1993/toto-todo-cli")
	}

	return sanitisedInput
}
