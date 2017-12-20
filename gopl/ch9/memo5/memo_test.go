package memo_test

import (
	"github.com/happyfire/gostudy/gopl/ch9/memo5"
	"github.com/happyfire/gostudy/gopl/ch9/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Concurrent(t, m)
}
