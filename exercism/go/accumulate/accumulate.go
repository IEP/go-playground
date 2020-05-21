package accumulate

// Accumulate will convert each collection of string given operator function
func Accumulate(str []string, converter func(string) string) []string {
	converted := make([]string, len(str))
	for index, value := range str {
		converted[index] = converter(value)
	}
	return converted
}
