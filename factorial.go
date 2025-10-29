package main


func Factorial(a float64) float64 {
	if a > 1 {
		return a * Factorial(a - 1)
	} else {
		return a
	}
}

func FactorialIter(a float64) float64 {
	var sum float64 = 1
	for i := 1.0; i <= a; i++ {
		sum *= i
	}
	return sum
}
