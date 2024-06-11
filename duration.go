package safe

import "time"

type SafeDuration struct{ *safeNumber }

func (t *SafeDuration) Set(value time.Duration) {
	_ = t.execLocked(func(st *safeNumber) any { st.value = value; return nil })
}

func (t *SafeDuration) Get() time.Duration {
	return t.execLocked(func(st *safeNumber) any { return st.value }).(time.Duration)
}

func (f *SafeDuration) Add(v time.Duration) { f.Set(f.Get() + v) }
func (f *SafeDuration) Div(v time.Duration) { f.Set(f.Get() / v) }
func (f *SafeDuration) Mul(v time.Duration) { f.Set(f.Get() * v) }
func (f *SafeDuration) Inc()                { f.Add(1) }
func (f *SafeDuration) Dec()                { f.Add(-1) }

func NewSafeDuration[D Durations](initValue D) *SafeDuration {
	return &SafeDuration{safeNumber: newSafeNumber(time.Duration(initValue))}
}
