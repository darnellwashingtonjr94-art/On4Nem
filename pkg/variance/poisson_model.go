package variance

import "math"

func PoissonProbability(lambda float64, k int) float64 {
	// Poisson distribution modeling for low-scoring sports (Soccer goals, Hockey pucks)
	return (math.Pow(lambda, float64(k)) * math.Exp(-lambda)) / float64(factorial(k))
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}
