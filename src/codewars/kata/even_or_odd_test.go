package kata

import (
	"testing"
)

func TestEvenOrOdd(t *testing.T) {
	// tests := Bar(10)

	type pair struct {
		val int
		res string
	}

	var tests = []pair{
		{1, "Odd"},
		{2, "Even"},
	}

	for _, pair := range tests {
		v := EvenOrOdd(pair.val)
		if v != pair.res {
			t.Error("Failed")
		}
	}

}
