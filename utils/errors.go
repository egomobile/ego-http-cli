package utils

import (
	"log"
	"os"
)

// ExitForInvalidArg() outputs an error message
// for an invalid argument and exits with code 1
func ExitForInvalidArg(a ...interface{}) {
	args := append([]interface{}{"⚠️"}, a...)

	log.Println(args...)
	os.Exit(2)
}
