// Package gnum provides generic numerical functions.
package gnum

import (
	"fmt"
	"math"

	"golang.org/x/exp/constraints"
)

// Number is a constraint that contains comparable numbers.
type Number interface {
	constraints.Float | constraints.Integer
}

// Max returns the maximal value in the slice or zero if the slice is empty.
func Max[S ~[]N, N constraints.Ordered](s S) N {
	if len(s) == 0 {
		var zero N
		return zero
	}
	e := s[0]
	for _, v := range s[1:] {
		e = max(e, v)
	}
	return e
}

// Min returns the maximal value in the slice or zero if the slice is empty.
func Min[S ~[]N, N constraints.Ordered](s S) N {
	if len(s) == 0 {
		var zero N
		return zero
	}
	e := s[0]
	for _, v := range s[1:] {
		e = min(e, v)
	}
	return e
}

// ArgMax returns the index of the maximal value in the slice or -1 if the slice is empty.
func ArgMax[S ~[]E, E constraints.Ordered](s S) int {
	if len(s) == 0 {
		return -1
	}
	imax, max := 0, s[0]
	for i, v := range s {
		if v > max {
			imax, max = i, v
		}
	}
	return imax
}

// ArgMin returns the index of the minimal value in the slice or -1 if the slice is empty.
func ArgMin[S ~[]E, E constraints.Ordered](s S) int {
	if len(s) == 0 {
		return -1
	}
	imin, min := 0, s[0]
	for i, v := range s {
		if v < min {
			imin, min = i, v
		}
	}
	return imin
}

// Abs returns the absolute value of n.
func Abs[N Number](n N) N {
	if n < 0 {
		return -n
	}
	return n
}

// Diff returns the non-negative difference between a and b.
func Diff[N Number](a, b N) N {
	if a > b {
		return a - b
	}
	return b - a
}

// Sum returns the sum of the slice.
func Sum[S ~[]N, N Number](a S) N {
	var sum N
	for _, v := range a {
		sum += v
	}
	return sum
}

// Mean returns the average of the slice.
func Mean[S ~[]N, N Number](a S) float64 {
	return float64(Sum(a)) / float64(len(a))
}

// Cov returns the covariance of a and b.
func Cov[S ~[]N, N Number](a, b S) float64 {
	assertMatchingLengths(a, b)
	ma := Mean(a)
	mb := Mean(b)
	cov := 0.0
	for i := range a {
		cov += (float64(a[i]) - ma) * (float64(b[i]) - mb)
	}
	cov /= float64(len(a))
	return cov
}

// Var returns the variance of a.
func Var[S ~[]N, N Number](a S) float64 {
	return Cov(a, a)
}

// Std returns the standard deviation of a.
func Std[S ~[]N, N Number](a S) float64 {
	return math.Sqrt(Var(a))
}

// Corr returns the Pearson correlation between the a and b.
func Corr[S ~[]N, N Number](a, b S) float64 {
	return Cov(a, b) / Std(a) / Std(b)
}

// Entropy returns the Shannon-entropy of a.
// The elements in a don't have to sum up to 1.
func Entropy[S ~[]N, N Number](a S) float64 {
	sum := float64(Sum(a))
	result := 0.0
	for i, v := range a {
		if v < 0.0 {
			panic(fmt.Sprintf("negative value at position %d: %v",
				i, v))
		}
		if v == 0 {
			continue
		}
		p := float64(v) / sum
		result -= p * math.Log2(p)
	}
	return result
}

// Idiv divides a by b, rounded to the nearest integer.
func Idiv[T constraints.Integer](a, b T) T {
	return T(math.Round(float64(a) / float64(b)))
}
