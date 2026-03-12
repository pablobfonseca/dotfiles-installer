package utils

import (
	"fmt"
	"log"

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

// InfoMessage prints an information message. Supports fmt.Sprintf-style format args.
func InfoMessage(message string, args ...interface{}) {
	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}
	fmt.Printf("%v %s\n", emoji.Information, message)
}
