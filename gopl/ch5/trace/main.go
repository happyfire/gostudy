// The trace program uses defer to add entry/exit diagnostics to function.

package main

import "time"
import "log"

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")() //注意：这儿推迟执行的是trace返回的闭包，而不是trace本身，trace本身是立刻执行的
	// ... losts of work ...
	time.Sleep(3 * time.Second) // simulate slow operation by sleeping
}

func main() {
	bigSlowOperation()
}
