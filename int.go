package safe

type SafeInt struct{ *safeNumber }

func (i *SafeInt) Get() int {
	return i.execLocked(func(st *safeNumber) any { return st.value }).(int)
}
func (i *SafeInt) Add(v int) int { return i.Set(i.Get() + v).(int) }
func (i *SafeInt) Div(v int) int { return i.Set(i.Get() / v).(int) }
func (i *SafeInt) Mul(v int) int { return i.Set(i.Get() * v).(int) }
func (i *SafeInt) Inc() int      { return i.Add(1) }
func (i *SafeInt) Dec() int      { return i.Add(-1) }

func NewSafeInt[I Ints](initValue I) *SafeInt {
	return &SafeInt{safeNumber: newSafeNumber(int(initValue))}
}
