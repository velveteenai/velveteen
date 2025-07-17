package main

// ReverseString reverses the characters in a string.
func ReverseString(text string) string {
	runes := []rune(text)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// CountWords returns the number of words in a sentence.
func CountWords(sentence string) int {
	words := SplitWords(sentence)
	return len(words)
}

// SplitWords splits a sentence into words by whitespace.
func SplitWords(sentence string) []string {
	return Fields(sentence)
}

// Fields splits a string by whitespace (helper for CountWords).
func Fields(s string) []string {
	var fields []string
	field := ""
	for _, r := range s {
		if r == ' ' || r == '\t' || r == '\n' {
			if field != "" {
				fields = append(fields, field)
				field = ""
			}
		} else {
			field += string(r)
		}
	}
	if field != "" {
		fields = append(fields, field)
	}
	return fields
}

// CelsiusToFahrenheit converts Celsius temperature to Fahrenheit.
func CelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}
