// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1700/E
// https://codeforces.com/problemset/status/1700/problem/E
func Test_cf1700E(t *testing.T) {
	testCases := [][2]string{
		{
			`3 3
2 1 3
6 7 4
9 8 5`,
			`0`,
		},
		{
			`2 3
1 6 4
3 2 5`,
			`1 3`,
		},
		{
			`1 6
1 6 5 4 3 2`,
			`2`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1700E)
}