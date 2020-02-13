package econ

import "math"

// SinglePaymentCompoundAmount (F/P, i, n).
// Moves a single payment to n periods later in time.
func SinglePaymentCompoundAmount(i float64, n int) float64 {
	return math.Pow(1+i, float64(n))
}

// SinglePaymentPresentWorth (F/P, i, n).
// Moves a single payment to n periods earlier in time.
func SinglePaymentPresentWorth(i float64, n int) float64 {
	return 1 / math.Pow(1+i, float64(n))
}

// UniformSeriesSinkingFund (A/F, i, n).
// Takes a single payment and spreads into a uniform series
// over n earlier periods. The last payment in the series
// occurs at the same time as F.
func UniformSeriesSinkingFund(i float64, n int) float64 {
	return i / (math.Pow(1+i, float64(n)) - 1)
}

// UniformSeriesCompoundAmount (F/A, i, n).
// Takes a uniform series and moves it to a single value at the
// time of the last payment in the series.
func UniformSeriesCompoundAmount(i float64, n int) float64 {
	return (math.Pow(1+i, float64(n)) - 1) / i
}

// UniformSeriesCapitalRecovery (A/P, i, n).
// Takes a single payment and spreads it into a uniform series
// over n later periods. The first payment in the series occurs
// one period later than P.
func UniformSeriesCapitalRecovery(i float64, n int) float64 {
	return (i * math.Pow(1+i, float64(n))) /
		(math.Pow(1+i, float64(n)) - 1)
}

// UniformSeriesPresentWorth (P/A, i, n).
// Takes a uniform series and moves it to a single payment one
// period earlier than the first payment of the series.
func UniformSeriesPresentWorth(i float64, n int) float64 {
	return (math.Pow(1+i, float64(n)) - 1) /
		(i * math.Pow(1+i, float64(n)))
}

// ArithmeticGradientPresentWorth (P/G, i, n).
// Takes a arithmetic gradient series and moves it to a single payment
// two periods earlier than the first nonzero payment of the series.
func ArithmeticGradientPresentWorth(i float64, n int) float64 {
	return (math.Pow(1+i, float64(n)) - i*float64(n) - 1) /
		(math.Pow(i, 2) * math.Pow(1+i, float64(n)))
}

// ArithmeticGradientToUniformSeries (A/G, i, n).
// Takes a arithmetic gradient series and converts it to a uniform series.
// The two series cover the same interval, but the first payment of the
// gradient series is 0.
func ArithmeticGradientToUniformSeries(i float64, n int) float64 {
	return (math.Pow(1+i, float64(n)) - i*float64(n) - 1) /
		(i*math.Pow(1+i, float64(n)) - i)
}

// GeometricSeriesPresentWorth (P/A, g, i, n).
// Calculates the present worth of a geometric gradient.
func GeometricSeriesPresentWorth(g, i float64, n int) float64 {
	if i == g {
		return float64(n) / (1 + i)
	}
	return (1 - math.Pow((1+g)/(1+i), float64(n))) / (i - g)
}

// ContinuousCompoundingSinkingFund [A/F, r, n].
func ContinuousCompoundingSinkingFund(r float64, n int) float64 {
	return (math.Pow(math.E, r) - 1) /
		(math.Pow(math.E, r*float64(n)) - 1)
}

// ContinuousCompoundingCapitalRecovery [A/P, r, n].
func ContinuousCompoundingCapitalRecovery(r float64, n int) float64 {
	return (math.Pow(math.E, r*float64(n)) * (math.Pow(math.E, r) - 1)) /
		(math.Pow(math.E, r*float64(n)) - 1)
}

// ContinuousCompoundingSeriesCompoundAmount [F/A, r, n].
func ContinuousCompoundingSeriesCompoundAmount(r float64, n int) float64 {
	return (math.Pow(math.E, r*float64(n)) - 1) /
		(math.Pow(math.E, r) - 1)
}

// ContinuousCompoundingSeriesPresentWorth [P/A, r, n].
func ContinuousCompoundingSeriesPresentWorth(r float64, n int) float64 {
	return (math.Pow(math.E, r*float64(n)) - 1) /
		(math.Pow(math.E, r*float64(n)) * (math.Pow(math.E, r) - 1))
}
