package tools

func RemoveElementFromSlice(slice []string, position int) []string {
	returnedSlice := make([]string, 0, len(slice)-2)
	for pos, value := range slice {
		if pos == position {
			continue
		}
		returnedSlice = append(returnedSlice, value)
	}
	return returnedSlice
}
