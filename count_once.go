package count_once

import (
	"sync"
	"sync/atomic"
)

type QueueMutex struct {
	flag int32
	mtx  sync.Mutex
}

func NewQueueMutex() *QueueMutex {
	return &QueueMutex{
		flag: 0,
		mtx:  sync.Mutex{},
	}
}

//并发时单次执行f函数
func (m *QueueMutex) Do(f func()) {
	v := atomic.LoadInt32(&m.flag)

	m.mtx.Lock()
	defer m.mtx.Unlock()

	if atomic.LoadInt32(&m.flag) != v {
		return
	}

	if f != nil {
		f()
	}

	atomic.StoreInt32(&m.flag, v+1)
}
