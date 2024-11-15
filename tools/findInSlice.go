package tools

import "fmt"

func FindInSlice(list []string, toSearch string) (int, error) {
	for i, v := range list {
		if v == toSearch {
			return i, nil
		}
	}
	return 0, fmt.Errorf("Not found")
}
