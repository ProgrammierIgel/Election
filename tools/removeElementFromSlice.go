package tools

func RemoveElementFromSlice(slice []string, position int) []string {
	slice[position] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

// Other Way

// func RemoveIndex(s []string, index int) []string {
// 	return append(s[:index], s[index+1:]...)
// }
