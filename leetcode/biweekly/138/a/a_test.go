// Generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_a(t *testing.T) {
	if err := testutil.RunLeetCodeFuncWithFile(t, generateKey, "a.txt", 0); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/biweekly-contest-138/problems/find-the-key-of-the-numbers/
// https://leetcode.cn/problems/find-the-key-of-the-numbers/