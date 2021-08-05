package super_functions

func Alinka(x, y int) int {
	// Суммирует два числа и возвращает результат сложения
	// Паникует, если результат больше 100

	sum := x + y
	if sum > 100 {
		panic("AAAA! это число слишком большое!")
	}
	return sum
}
