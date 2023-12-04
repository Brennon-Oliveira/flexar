package utils

import (
	"fmt"
	"os"
)

func SemanticError(line int, message string) {
	fmt.Printf("(%s) SemanticError on line %d: %s\n", GetCurrentFile(), line, message)
	os.Exit(0)
}
