package tools

import "fmt"

func GetIndexOfElementInSlice(slice []string, element string) (int, error) {
	for index, value := range slice {
		if value == element {
			return index, nil
		}
	}
	return -1, fmt.Errorf("elemtent not found")
}
