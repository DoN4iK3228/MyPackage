package mymath

import "math"

// Sqrt возвращает квадратный корень числа
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

// Ceil возвращает наименьшее целое число, не меньшее x
func Ceil(x float64) float64 {
	return math.Ceil(x)
}

// Floor возвращает наибольшее целое число, не большее x
func Floor(x float64) float64 {
	return math.Floor(x)
}

// Pow возводит число x в степень y
func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

// Max возвращает большее из двух чисел
func Max(x, y float64) float64 {
	return math.Max(x, y)
}

// Min возвращает меньшее из двух чисел
func Min(x, y float64) float64 {
	return math.Min(x, y)
}

func Abs(x float64) float64 {
    return math.Abs(x)
}

// Acos возвращает арккосинус числа
func Acos(x float64) float64 {
    return math.Acos(x)
}

// Acosh возвращает гиперболический арккосинус числа
func Acosh(x float64) float64 {
    return math.Acosh(x)
}

// Asin возвращает арксинус числа
func Asin(x float64) float64 {
    return math.Asin(x)
}

func Yn(n int, x float64) float64 {
	return math.Yn(x)
}
