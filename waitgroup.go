
package main

import (
	"sync/atomic"

)



func (wg *WaitGroupCount) Add(delta int) {
    atomic.AddInt64(&wg.count, int64(delta))
    wg.WaitGroup.Add(delta)
}

func (wg *WaitGroupCount) Done() {
    atomic.AddInt64(&wg.count, -1)
    wg.WaitGroup.Done()
}

func (wg *WaitGroupCount) GetCount() int {
    return int(atomic.LoadInt64(&wg.count))
}

// Wait() promoted from the embedded field