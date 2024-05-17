package helpers

func IsArrayContains(arr []string, str string) bool {
	for _, item := range arr {
		if item == str {
			return true
		}
	}
	return false
}

func ConcatString(str1 string, str2 string) string {
	return str1 + str2
}
