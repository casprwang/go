package kata

import (
	"testing"
)

func TestRepeat(t *testing.T) {

	type input struct {
		rep int
		str string
	}

	type pair struct {
		input
		res string
	}

	var tests = []pair{
		{input{1, "Odd"}, "Odd"},
		{input{2, "Even"}, "EvenEven"},
	}

	// tests := Bar(10)
	for _, pair := range tests {
		v := RepeatStr(pair.input.rep, pair.input.str)
		if v != pair.res {
			t.Error("Failed")
		}
	}

}
