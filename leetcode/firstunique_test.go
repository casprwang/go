package leetcode

import (
	"testing"
)

var tables = []struct {
	in  string
	out int
}{
	{"leetcode", 0},
	{"loveleetcode", 2},
}

func Test(t *testing.T) {

	for _, table := range tables {
		res := firstUniqChar(table.in)

		if res != table.out {
			t.Errorf("Case: %v, got: %v, need: %v", table.in, res, table.out)
		}
	}
}
