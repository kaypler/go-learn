package memo_test

import (
	"testing"

	"go-learn/ch9/memotest"
	memo "go-learn/go-learn/ch9/memo2"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}
