package play

import (
	"testing"
)

func TestNumber(t *testing.T) {
	// tests := Bar(10)
	tests := [3]int{10, 10, 10}

	for _, val := range tests {
		if val != int(Bar(10)) {
			t.Error("Failed")
		}
	}

	// if 10 != tests {
	// 	t.Error("Failed")
	// }
}

func TestInteger(t *testing.T) {
	tests := [4]int{1, 2, 4, 5}

	for _, val := range tests {
		if val > 10 {
			t.Error(">10")
		}
	}
}
