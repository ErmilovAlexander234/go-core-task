package main

import "testing"

var (
	decimalInt int       = 42
	octalInt   int       = 0755
	hexInt     int       = 0xFF
	floatNum   float64   = 3.14159
	str        string    = "Golang"
	boolVal    bool      = true
	complexNum complex64 = 1 + 2i
)

func TestGetType(t *testing.T) {
	tests := map[interface{}]string{
		decimalInt: "int",
		octalInt:   "int",
		hexInt:     "int",
		floatNum:   "float64",
		str:        "string",
		boolVal:    "bool",
		complexNum: "complex64",
	}

	for v, expected := range tests {
		if GetType(v) != expected {
			t.Errorf("Expected %s", expected)
		}
	}
}

func TestConvertToString(t *testing.T) {
	if ConvertToString(decimalInt) != "42" {
		t.Error("Decimal conversion failed")
	}

	if ConvertToString(boolVal) != "true" {
		t.Error("Bool conversion failed")
	}
}

func TestRunesAndHash(t *testing.T) {
	combined := CombineStrings(
		decimalInt,
		octalInt,
		hexInt,
		floatNum,
		str,
		boolVal,
		complexNum,
	)

	runes := StringToRunes(combined)

	if len(runes) == 0 {
		t.Error("Empty runes")
	}

	hash1 := HashRunesWithSalt(runes)
	hash2 := HashRunesWithSalt(runes)

	if hash1 != hash2 {
		t.Error("Not stable hash")
	}
}
