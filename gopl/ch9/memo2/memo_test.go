package memo_test

import (
	"github.com/happyfire/gostudy/gopl/ch9/memo2"
	"github.com/happyfire/gostudy/gopl/ch9/memotest"
	"testing"
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
