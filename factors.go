package econ

import "math"

// FP is the Single Payment Compound Amount Factor (F/P, i, n).
// Moves a single payment to n periods later in time.
func FP(i float64, n int) float64 {
	return math.Pow(1+i, float64(n))
}

// PF is the Single Payment Present Worth Factor (F/P, i, n).
// Moves a single payment to n periods earlier in time.
func PF(i float64, n int) float64 {
	return 1 / math.Pow(1+i, float64(n))
}

// AF is the Uniform Series Sinking Fund Factor (A/F, i, n).
// Takes a single payment and spreads into a uniform series
// over n earlier periods. The last payment in the series
// occurs at the same time as F.
func AF(i float64, n int) float64 {
	return i / (math.Pow(1+i, float64(n)) - 1)
}

// FA is the Uniform Series Compound Amount Factor (F/A, i, n).
// Takes a uniform series and moves it to a single value at the
// time of the last payment in the series.
func FA(i float64, n int) float64 {
	return (math.Pow(1+i, float64(n)) - 1) / i
}

// AP is the Uniform Series Capital Recovery Factor (A/P, i, n).
// Takes a single payment and spreads it into a uniform series
// over n later periods. The first payment in the series occurs
// one period later than P.
func AP(i float64, n int) float64 {
	return (i * math.Pow(1+i, float64(n))) /
		(math.Pow(1+i, float64(n)) - 1)
}

// PA is the Uniform Series Present Worth Factor (P/A, i, n).
// Takes a uniform series and moves it to a single payment one
// period earlier than the first payment of the series.
func PA(i float64, n int) float64 {
	return (math.Pow(1+i, float64(n)) - 1) /
		(i * math.Pow(1+i, float64(n)))
}

// PG is the Arithmetic Gradient Present Worth Factor (P/G, i, n).
// Takes a arithmetic gradient series and moves it to a single payment
// two periods earlier than the first nonzero payment of the series.
func PG(i float64, n int) float64 {
	return (math.Pow(1+i, float64(n)) - i*float64(n) - 1) /
		(math.Pow(i, 2) * math.Pow(1+i, float64(n)))
}

// AG is the Arithmetic Gradient to Uniform Series Factor (A/G, i, n).
// Takes a arithmetic gradient series and converts it to a uniform series.
// The two series cover the same interval, but the first payment of the
// gradient series is 0.
func AG(i float64, n int) float64 {
	return (math.Pow(1+i, float64(n)) - i*float64(n) - 1) /
		(i*math.Pow(1+i, float64(n)) - i)
}
