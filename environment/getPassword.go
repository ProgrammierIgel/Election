package environment

import (
	"fmt"
	"os"
)

func Password() (string, error) {
	password := os.Getenv("PASSWORD")

	if password == "" {
		return "", fmt.Errorf("NoPassword")
	}

	return password, nil
}
