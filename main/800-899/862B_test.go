// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/862/B
// https://codeforces.com/problemset/status/862/problem/B
func Test_cf862B(t *testing.T) {
	testCases := [][2]string{
		{
			`3
1 2
1 3`,
			`0`,
		},
		{
			`5
1 2
2 3
3 4
4 5`,
			`2`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf862B)
}