package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
)

func GetType(i interface{}) string {
	return reflect.TypeOf(i).String()
}

func ConvertToString(i interface{}) string {
	return fmt.Sprintf("%v", i)
}

func CombineStrings(values ...interface{}) string {
	result := ""
	for _, v := range values {
		result += ConvertToString(v)
	}
	return result
}

func StringToRunes(s string) []rune {
	return []rune(s)
}

func HashRunesWithSalt(r []rune) string {
	s := string(r)

	mid := len(s) / 2
	salted := s[:mid] + "go-2024" + s[mid:]

	hash := sha256.Sum256([]byte(salted))

	return fmt.Sprintf("%x", hash)
}

func main() {

	var (
		decimalInt int       = 42
		octalInt   int       = 052
		hexInt     int       = 0x2A
		floatNum   float64   = 3.14159
		str        string    = "Golang"
		boolVal    bool      = true
		complexNum complex64 = 1 + 2i
	)

	fmt.Println("Types:")

	fmt.Println(GetType(decimalInt))
	fmt.Println(GetType(octalInt))
	fmt.Println(GetType(hexInt))
	fmt.Println(GetType(floatNum))
	fmt.Println(GetType(str))
	fmt.Println(GetType(boolVal))
	fmt.Println(GetType(complexNum))

	combined := CombineStrings(
		decimalInt,
		octalInt,
		hexInt,
		floatNum,
		str,
		boolVal,
		complexNum,
	)

	fmt.Println("\nCombinedString:")
	fmt.Println(combined)

	runes := StringToRunes(combined)

	hash := HashRunesWithSalt(runes)

	fmt.Println("\nSHA256 Hash:")
	fmt.Println(hash)
}
