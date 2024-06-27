package gyro

import (
	"fmt"
	"math"
	"math/big"
	"strconv"

	"github.com/profe-ajedrez/gyro/i128"
)

type Gyro struct {
	coeff i128.I128
	exp   int32
}

// New returns a new Gyro with the given coefficient and exponent.
// if exp greater than maxScale, it will be clamped to maxScale.
// if exp is positive, will multiply the coefficient by 10^exp.
func New(coeff i128.I128, exp int32) Gyro {

	if exp == 0 {
		return Gyro{coeff, 0}
	}

	if exp < 0 {
		if abs(exp) > maxScale {
			exp = -maxScale
			coeff = coeff.Quo64(pow10max16(-exp))
		}
	}

	if exp > 0 {
		exp = min(exp, maxScale)
		coeff = coeff.Mul64(pow10max16(exp))
		exp = 0
	}

	return Gyro{coeff, exp}
}

// NewFromInt64Raw returns a new Gyro with the given coefficient and exponent.
func NewFromInt64Raw(coeff int64, exp int32) Gyro {
	return New(i128.I128From64(coeff), exp)
}

// NewFromInt32Raw returns a new Gyro with the given int32 coefficient and exponent.
func NewFromInt32Raw(coeff int32, exp int32) Gyro {
	return New(i128.I128From32(coeff), exp)
}

func NewFromInt64(coeff int64) Gyro {
	return New(i128.I128From64(coeff), 0)
}

// NewFromString returns a string representation of a Gyro Decimal
func NewFromString(s string) (Gyro, error) {
	if s == "" {
		return Gyro{}, NewGyroInvalidErr("couldnt convert an empty space in a decimal value")
	}

	neg := s[0] == '-'
	start := 0

	if neg {

		if len(s) == 1 {
			return Gyro{}, NewGyroInvalidErr("couldnt convert a negative sign in a decimal value")
		}

		start = 1
	}

	dIndex := 0
	var i int
	for i = start; i < len(s); i++ {
		if s[i] == '.' {
			if dIndex != 0 {
				return Gyro{}, NewGyroInvalidErr("only decimal strings are supported. " + s + " has more than 1 decimal point")
			}
			dIndex = i
		}
	}

	var exp int32 = int32(i) - int32(dIndex) - 1

	if dIndex == 0 {
		exp = 0
	}

	if exp > maxScale {
		exp = maxScale
		i = maxScale + dIndex
	}

	if dIndex == 0 {

		u64, err := strconv.ParseUint(s[start:i], 10, 64)

		if err != nil {
			return newFromString(s, dIndex, i, exp)
		}

		d := i128.I128FromU64(u64)

		if neg {
			d = d.Mul64(-1)
		}
		return Gyro{coeff: d, exp: 0}, nil
	}

	c, err := strconv.ParseUint(s[start:dIndex], 10, 64)

	if err != nil {
		return newFromString(s, dIndex, i, exp)
	}

	e, err := strconv.ParseUint(s[dIndex+1:i], 10, 64)

	if err != nil {
		return newFromString(s, dIndex, i, exp)
	}

	d := i128.I128FromU64(c)
	d = d.Mul64(pow10max16(int32(len(s[dIndex+1 : i]))))
	ex := i128.I128FromU64(e)
	d = d.Add(ex)

	if neg {
		d = d.Mul64(-1)
	}
	return Gyro{coeff: d, exp: -exp}, nil
}

func newFromString(s string, dIndex int, i int, exp int32) (Gyro, error) {

	if dIndex == 0 {
		s = s[:i]
	} else {
		s = s[:dIndex] + s[dIndex+1:i]
	}

	coeff, _, err := i128.I128FromString(s)

	if err != nil {
		return Gyro{}, NewGyroInvalidErr("couldnt convert " + s + " in a decimal value")
	}

	return Gyro{coeff: coeff, exp: -exp}, nil
}

func (g Gyro) String() string {
	if g.exp == 0 {
		return g.coeff.String()
	}

	ex := abs(g.exp)
	coeff := g.coeff.Quo64(pow10max16(ex))
	exp := g.coeff.Sub(coeff.Mul64(pow10max16(ex))).Abs()
	return coeff.String() + "." + exp.String()
}

func (g Gyro) Int64() int64 {
	return g.coeff.AsInt64()
}

func (g Gyro) Float64() float64 {
	coeff := g.coeff.AsFloat64()
	exp := int(g.exp)

	if exp >= 0 {
		return coeff * pow10Float64(exp)
	} else {
		return coeff / pow10Float64(-exp)
	}
}

func pow10Float64(exp int) float64 {
	if exp < 16 {
		return float64(pow10tab[exp])
	}
	return math.Pow10(exp)
}

func (g Gyro) IsFloat64Exact() bool {
	f64 := g.Float64()
	bf := new(big.Float).SetInt64(g.coeff.AsInt64())
	exp := new(big.Float).SetInt64(int64(g.exp))
	exp.Neg(exp)
	expInt, _ := exp.Int64()
	bf.SetMantExp(bf, int(expInt))
	bf64, _ := bf.Float64()
	diff := math.Abs(f64 - bf64)

	// Adjust the tolerance as needed
	const tolerance = 1e-15

	return diff <= tolerance
}

func (g Gyro) Add(g2 Gyro) Gyro {
	g, g2 = normalize(g, g2)

	i128 := g.coeff.Add(g2.coeff)

	return Gyro{i128, g.exp}
}

func (g Gyro) Sub(g2 Gyro) Gyro {
	g, g2 = normalize(g, g2)

	i128 := g.coeff.Sub(g2.coeff)

	return Gyro{i128, g.exp}
}

func (g Gyro) Mul(g2 Gyro) Gyro {

	expInt64 := int64(g.exp) + int64(g2.exp)
	if expInt64 > math.MaxInt32 || expInt64 < math.MinInt32 {
		panic(fmt.Sprintf("exponent %v overflows an int32!", expInt64))
	}

	d3Value := g.coeff.Mul(g2.coeff)

	g3 := Gyro{
		coeff: d3Value,
		exp:   int32(expInt64),
	}

	if expInt64 > 0 {
		return g3.rescale(min(int32(expInt64), maxScale))
	}
	return g3.rescale(max(int32(expInt64), -maxScale))
}

// func (g Gyro) Div(g2 Gyro) (Gyro, error) {

// }

func (g Gyro) DivRound(g2 Gyro, scale int32) Gyro {
	q, r := g.QuoRem(g2, scale)
	rv2 := r.coeff.Abs().Mul64(2)
	r2 := Gyro{rv2, r.exp + scale}
	c := r2.Cmp(g2.Abs())

	if c < 0 {
		return q
	}

	if g.coeff.Sign()*g.coeff.Sign() < 0 {
		return Gyro{i128.I128From64(1), -scale}
	}

	i, _ := i128.I128FromFloat64(1)
	return q.Add(New(i, -scale))

}

func (g Gyro) QuoRem(g2 Gyro, precision int32) (Gyro, Gyro) {

	scale := -precision
	e := int64(g.exp) - int64(g2.exp) - int64(scale)

	if e > math.MaxInt32 || e < math.MinInt32 {
		panic("overflow in decimal QuoRem")
	}

	var aa, bb, expo i128.I128
	var scalerest int32

	if e < 0 {
		aa = g.coeff
		expo = i128.I128From64(pow10max16(-e))
		bb = tenInt.Mul(expo)
		scalerest = g.exp
	} else {
		expo = i128.I128From64(pow10max16(-e))
		aa = g.coeff.Mul(expo)
		bb = g2.coeff
		scalerest = int32(scale) + g2.exp
	}

	q, r := aa.QuoRem(bb)

	dq := Gyro{q, scale}
	dr := Gyro{r, scalerest}

	return dq, dr

}

func (g Gyro) Equal(other Gyro) bool {
	// Normalize the coefficients and exponents
	g1, g2 := normalize(g, other)

	// Compare the normalized coefficients and exponents
	return g1.coeff.Equal(g2.coeff) && g1.exp == g2.exp
}

func (g Gyro) Abs() Gyro {
	if g.coeff.Sign() < 0 {
		return Gyro{coeff: g.coeff.Mul64(-1), exp: g.exp}
	}
	return g
}

// Cmp compares the numbers represented by d and d2 and returns:
//
//	-1 if d <  d2
//	 0 if d == d2
//	+1 if d >  d2
func (g Gyro) Cmp(g2 Gyro) int {

	if g.exp == g2.exp {
		return g.coeff.Cmp(g2.coeff)
	}

	rd, rd2 := normalize(g, g2)

	return rd.coeff.Cmp(rd2.coeff)
}

func normalize(g1, g2 Gyro) (Gyro, Gyro) {
	if g1.exp < g2.exp {
		return g1, g2.rescale(g1.exp)
	} else if g1.exp > g2.exp {
		return g1.rescale(g2.exp), g2
	}

	return g1, g2
}

func (g Gyro) rescale(exp int32) Gyro {

	if g.exp == exp {
		return Gyro{
			g.coeff,
			g.exp,
		}
	}

	diff := abs(exp - g.exp)
	value := g.coeff

	expScale := i128.I128FromRaw(0, uint64(pow10max16(diff)))

	if exp > g.exp {
		value = value.Quo(expScale)
	} else if exp < g.exp {
		value = value.Mul(expScale)
	}

	return Gyro{
		coeff: value,
		exp:   exp,
	}
}
