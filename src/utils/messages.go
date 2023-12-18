package utils

import (
	"fmt"
	"log"
	"strings"

	"github.com/enescakir/emoji"
)

// ErrorMessage prints an error message and exits the program
func ErrorMessage(message string, err error) {
	log.Fatalf("%v %s: %v", emoji.CrossMark, message, err)
}

// SuccessMessage prints a success message
func SuccessMessage(message string) {
	fmt.Printf("%v %s\n", emoji.CheckMark, message)
}

// SkipMessage prints a skipping message
func SkipMessage(message string) {
	fmt.Printf("%v %s, skipping...\n", emoji.CheckMark, message)
}

// InfoMessage prints an information message
func InfoMessage(message string, args ...interface{}) {
	var stringArgs []string

	for _, arg := range args {
		stringArgs = append(stringArgs, fmt.Sprintf("%v", arg))
	}

	fullMessage := message
	if len(args) > 0 {
		fullMessage += ": " + strings.Join(stringArgs, ", ")
	}

	fmt.Printf("%v %s\n", emoji.Information, fullMessage)
}
