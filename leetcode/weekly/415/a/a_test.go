// Generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_a(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, getSneakyNumbers, "a.txt", 0); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/weekly-contest-415/problems/the-two-sneaky-numbers-of-digitville/
// https://leetcode.cn/problems/the-two-sneaky-numbers-of-digitville/