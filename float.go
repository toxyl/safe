package safe

type SafeFloat struct{ *safeNumber }

func (f *SafeFloat) Get() float64 {
	return f.execLocked(func(st *safeNumber) any { return st.value }).(float64)
}

func (f *SafeFloat) Add(v float64) float64 { return f.Set(f.Get() + v).(float64) }
func (f *SafeFloat) Div(v float64) float64 { return f.Set(f.Get() / v).(float64) }
func (f *SafeFloat) Mul(v float64) float64 { return f.Set(f.Get() * v).(float64) }
func (f *SafeFloat) Inc() float64          { return f.Add(1) }
func (f *SafeFloat) Dec() float64          { return f.Add(-1) }

func NewSafeFloat[F Floats](initValue F) *SafeFloat {
	return &SafeFloat{safeNumber: newSafeNumber(float64(initValue))}
}
