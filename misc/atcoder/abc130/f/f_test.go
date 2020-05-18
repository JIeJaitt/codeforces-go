// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [f]")
	testCases := [][2]string{
		{
			`2
0 3 D
3 0 L`,
			`0`,
		},
		{
			`5
-7 -10 U
7 -6 U
-8 7 D
-3 3 D
0 -6 R`,
			`97.5`,
		},
		{
			`20
6 -10 R
-4 -9 U
9 6 D
-3 -2 R
0 7 D
4 5 D
10 -10 U
-1 -8 U
10 -6 D
8 -5 U
6 4 D
0 3 D
7 9 R
9 -4 R
3 10 D
1 9 U
1 -6 U
9 -8 R
6 7 D
7 -3 D`,
			`273`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
	//testutil.AssertEqualRunResults(t, testCases, 0, runAC, run)
}
// https://atcoder.jp/contests/abc130/tasks/abc130_f
