//并发安全的函数cache，使用Mutex实现
//由于用来锁，此版本性能很低
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
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	memo.mu.Unlock()
	return res.value, res.err
}
