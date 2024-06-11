package safe

type SafeUint struct{ *safeNumber }

func (u *SafeUint) Get() uint {
	return u.execLocked(func(st *safeNumber) any { return st.value }).(uint)
}

func (u *SafeUint) Add(v int) uint {
	old := u.Get()
	r := int(old) + v
	if r < 0 {
		return u.Set(uint(0)).(uint)
	}
	return u.Set(uint(r)).(uint)
}
func (u *SafeUint) Div(v uint) uint { return u.Set(u.Get() / v).(uint) }
func (u *SafeUint) Mul(v uint) uint { return u.Set(u.Get() * v).(uint) }
func (u *SafeUint) Inc() uint       { return u.Add(1) }
func (u *SafeUint) Dec() uint       { return u.Add(-1) }

func NewSafeUint[U Uints](initValue U) *SafeUint {
	i := &SafeUint{safeNumber: newSafeNumber(uint(initValue))}
	return i
}
