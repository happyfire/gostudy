//memo2的改进版，将临界区改成两个，分别是获取cache和更新cache,性能有所提升，但还是有问题：
//一些url被获取多次，多个goroutine一起查询cache，发现没有值，然后一起调用f，得到结果后都会去更新map
//结果会互相覆盖

package memo

import "sync"

type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]result
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// Get is concurrency-safe
func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()

	if !ok {
		res.value, res.err = memo.f(key)

		//在两个临界区中间，几个goroutine可以竞争计算f(key)
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}

	return res.value, res.err
}
