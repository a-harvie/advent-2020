package errors

import "log"

// Check will blow your program up for you
func Check(err error) {
	if err != nil {
		log.Fatal("Oh noes", err)
	}
}
