package gyro

import (
	"github.com/profe-ajedrez/gyro/i128"
	"golang.org/x/exp/constraints"
)

const (
	maxScale = 16
)

var tenInt = i128.I128FromRaw(0, 10)

// pow10tab stores the pre-computed values 10**i for i < 32.
var pow10tab = [...]int64{
	1e00, 1e01, 1e02, 1e03, 1e04, 1e05, 1e06, 1e07, 1e08, 1e09,
	1e10, 1e11, 1e12, 1e13, 1e14, 1e15, 1e16, 1e17, 1e18,
}

func NewZero() Gyro {
	return Gyro{coeff: i128.I128{}, exp: 0}
}

func NewHundred() Gyro {
	return Gyro{coeff: i128.I128FromRaw(0, 100), exp: 0}
}

func NewTen() Gyro {
	return Gyro{coeff: tenInt, exp: 0}
}

func NewOne() Gyro {
	return Gyro{coeff: i128.I128FromRaw(0, 1), exp: 0}
}

func NewMinusOne() Gyro {
	return Gyro{coeff: i128.MustI128FromFloat64(-1), exp: 0}
}

func pow10max16[T constraints.Integer](v T) int64 {

	n := abs(v)

	if n > 16 {
		n = 16
	}

	if n == 0 {
		return 1
	}

	if n == 1 {
		return 10
	}

	return pow10tab[n]
}

func abs[T constraints.Integer](n T) int64 {
	if n < 0 {
		return int64(-n)
	}

	return int64(n)
}
