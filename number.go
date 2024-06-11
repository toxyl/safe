package safe

import "sync"

type safeNumber struct {
	mu    *sync.Mutex
	value any
}

func (t *safeNumber) execLocked(fn func(st *safeNumber) any) any {
	t.mu.Lock()
	defer t.mu.Unlock()
	return fn(t)
}

func (t *safeNumber) Set(value any) any {
	return t.execLocked(func(st *safeNumber) any { st.value = value; return value })
}

func (t *safeNumber) Get() any {
	return t.execLocked(func(st *safeNumber) any { return st.value })
}

func newSafeNumber[N Number](initValue N) *safeNumber {
	t := &safeNumber{
		mu:    &sync.Mutex{},
		value: initValue,
	}
	return t
}
