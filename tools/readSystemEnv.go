package tools

import (
	"fmt"
	"os"
)

// ReadSystemEnv reads a system environment variable by its name and returns its value.
// If the variable does not exist, it returns an empty string.
func ReadSystemEnv(varName string) string {
	value, exists := os.LookupEnv(varName)
	if !exists {
		fmt.Printf("Environment variable %s does not exist\n", varName)
		return ""
	}
	return value
}
