package main

import "strings"

func FirstLetterToUpperCase(x string) string {
	return strings.ToUpper(x[:1]) + strings.ToLower(x[1:])
}
