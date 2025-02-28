package safe

import "sync"

type SafeBool struct{ *safeNumber }

func (i *SafeBool) Get() bool {
	return i.execLocked(func(st *safeNumber) any { return st.value }).(bool)
}
func (i *SafeBool) Enable() bool  { return i.Set(true).(bool) }
func (i *SafeBool) Disable() bool { return i.Set(false).(bool) }
func (i *SafeBool) Toggle() bool  { return i.Set(!i.Get()).(bool) }

func NewSafeBool(initValue bool) *SafeBool {
	return &SafeBool{
		safeNumber: &safeNumber{
			mu:    &sync.Mutex{},
			value: initValue,
		},
	}
}
