package gothon

import (
	"log"
)

func Error(line int, message string) {
	report(line, "", message)
}

// Report an error at a specific line.
// if where is empty, the error is reported at the line.
func report(line int, where string, message string) {
	if where == "" {
		log.Fatalf("[line %d] Error: %s\n", line, message)
	} else {
		log.Fatalf("[line %d] Error at '%s': %s\n", line, where, message)
	}
}
