package super_functions

type SumObj struct {
	X         int
	Y         int
	Z         int
	AfterCalc bool
}

//go:generate minimock -i CanSum -o . -s _mock.go
type CanSum interface {
	Sum(x, y int) int
}

func (so *SumObj) Sum(x, y int) int {
	// Суммирует два числа и возвращает результат сложения
	// Паникует, если результат больше 100

	so.X = x
	so.Y = y
	so.Z = so.X + so.Y
	if so.Z > 100 {
		panic("AAAA! это число слишком большое!")
	}
	so.AfterCalc = true

	return so.Z
}
