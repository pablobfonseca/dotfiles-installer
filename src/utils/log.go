package utils

import (
	"fmt"
	"log"

	"github.com/enescakir/emoji"
)

func ErrorMessage(message string, err error) {
	log.Fatalf("%v %s: %v", emoji.CrossMark, message, err)
}

func SkipMessage(message string) {
	fmt.Printf("%v %s, skipping...\n", emoji.CheckMark, message)
}

func InfoMessage(message string) {
	fmt.Printf("%v %s\n", emoji.Information, message)
}
