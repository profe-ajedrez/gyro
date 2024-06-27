package gyro

import (
	"testing"

	"github.com/profe-ajedrez/gyro/i128"
)

func BenchmarkNew(b *testing.B) {
	testCases := []struct {
		name  string
		coeff i128.I128
		exp   int32
	}{
		{
			name:  "Zero coeff and exp",
			coeff: i128.I128{},
			exp:   0,
		},
		{
			name:  "Positive coeff and exp",
			coeff: i128.I128FromInt(123456789),
			exp:   5,
		},
		{
			name:  "Negative coeff and exp",
			coeff: i128.I128FromInt(-987654321),
			exp:   -5,
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = New(tc.coeff, tc.exp)
			}
		})
	}
}

func BenchmarkGyroString(b *testing.B) {
	testCases := []struct {
		name  string
		coeff i128.I128
		exp   int32
	}{
		{
			name:  "Zero coeff and exp",
			coeff: i128.I128{},
			exp:   0,
		},
		{
			name:  "Positive coeff and zero exp",
			coeff: i128.I128FromInt(12345),
			exp:   0,
		},
		{
			name:  "Negative coeff and zero exp",
			coeff: i128.I128FromInt(-67890),
			exp:   0,
		},
		{
			name:  "Positive coeff and positive exp",
			coeff: i128.I128FromInt(12345),
			exp:   3,
		},
		{
			name:  "Negative coeff and positive exp",
			coeff: i128.I128FromInt(-67890),
			exp:   5,
		},
		{
			name:  "Positive coeff and negative exp",
			coeff: i128.I128FromInt(12345),
			exp:   -2,
		},
		{
			name:  "Negative coeff and negative exp",
			coeff: i128.I128FromInt(-67890),
			exp:   -4,
		},
	}

	b.ResetTimer()
	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			gyro := New(tc.coeff, tc.exp)
			for i := 0; i < b.N; i++ {
				_ = gyro.String()
			}
		})
	}
}

func BenchmarkNewFromString(b *testing.B) {
	testCases := []struct {
		name  string
		input string
	}{
		{
			name:  "Valid integer",
			input: "12345",
		},
		{
			name:  "Valid negative integer",
			input: "-67890",
		},
		{
			name:  "Valid decimal with positive exponent",
			input: "12.345",
		},
		{
			name:  "Valid decimal with negative exponent",
			input: "67890.123",
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = NewFromString(tc.input)
			}
		})
	}
}

func BenchmarkGyroAdd(b *testing.B) {
	testCases := []struct {
		name string
		g1   Gyro
		g2   Gyro
	}{
		{
			name: "Zero and zero",
			g1:   New(i128.I128{}, 0),
			g2:   New(i128.I128{}, 0),
		},
		{
			name: "Positive and zero",
			g1:   NewFromInt64Raw(123, 0),
			g2:   New(i128.I128{}, 0),
		},
		{
			name: "Negative and zero",
			g1:   NewFromInt64Raw(-456, 0),
			g2:   New(i128.I128{}, 0),
		},
		{
			name: "Positive and positive",
			g1:   NewFromInt64Raw(123, 2),
			g2:   NewFromInt64Raw(456, 2),
		},
		{
			name: "Negative and negative",
			g1:   NewFromInt64Raw(-123, 2),
			g2:   NewFromInt64Raw(-456, 2),
		},
		{
			name: "Positive and negative",
			g1:   NewFromInt64Raw(123, 2),
			g2:   NewFromInt64Raw(-456, 2),
		},
		{
			name: "Negative and positive",
			g1:   NewFromInt64Raw(-123, 2),
			g2:   NewFromInt64Raw(456, 2),
		},
		{
			name: "Different exponents",
			g1:   NewFromInt64Raw(123, 2),
			g2:   NewFromInt64Raw(456, 4),
		},
		{
			name: "Large values",
			g1:   func() Gyro { g, _ := NewFromString("123456789012345678901"); return g }(),
			g2:   func() Gyro { g, _ := NewFromString("987654321098765432109"); return g }(),
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tc.g1.Add(tc.g2)
			}
		})
	}
}

func BenchmarkGyroSub(b *testing.B) {

	testCases := []struct {
		name string
		g1   Gyro
		g2   Gyro
	}{
		{
			name: "Zero minus zero",
			g1:   New(i128.I128{}, 0),
			g2:   New(i128.I128{}, 0),
		},
		{
			name: "Positive minus zero",
			g1:   NewFromInt64Raw(123, 0),
			g2:   New(i128.I128{}, 0),
		},
		{
			name: "Negative minus zero",
			g1:   NewFromInt64Raw(-456, 0),
			g2:   New(i128.I128{}, 0),
		},
		{
			name: "Positive minus positive (larger - smaller)",
			g1:   NewFromInt64Raw(456, 2),
			g2:   NewFromInt64Raw(123, 2),
		},
		{
			name: "Positive minus positive (smaller - larger)",
			g1:   NewFromInt64Raw(123, 2),
			g2:   NewFromInt64Raw(456, 2),
		},
		{
			name: "Negative minus negative (larger - smaller)",
			g1:   NewFromInt64Raw(-123, 2),
			g2:   NewFromInt64Raw(-456, 2),
		},
		{
			name: "Negative minus negative (smaller - larger)",
			g1:   NewFromInt64Raw(-456, 2),
			g2:   NewFromInt64Raw(-123, 2),
		},
		{
			name: "Positive minus negative",
			g1:   NewFromInt64Raw(123, 2),
			g2:   NewFromInt64Raw(-456, 2),
		},
		{
			name: "Negative minus positive",
			g1:   NewFromInt64Raw(-123, 2),
			g2:   NewFromInt64Raw(456, 2),
		},
		{
			name: "Different exponents",
			g1:   NewFromInt64Raw(123, 2),
			g2:   NewFromInt64Raw(456, 4),
		},
		{
			name: "Large values",
			g1:   func() Gyro { g, _ := NewFromString("123456789012345678901"); return g }(),
			g2:   func() Gyro { g, _ := NewFromString("987654321098765432109"); return g }(),
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tc.g1.Sub(tc.g2)
			}
		})
	}
}

func BenchmarkDivRound(b *testing.B) {

	tests := []struct {
		name      string
		g1        Gyro
		g2        Gyro
		precision int32
		expected  Gyro
	}{
		{
			"[ 0] -10.10112212 / 2.304 w 16 exp",
			func() Gyro { g, _ := NewFromString("-10.10112212"); return g }(),
			func() Gyro { g, _ := NewFromString("2.304"); return g }(), 1,
			func() Gyro { g, _ := NewFromString("-4.3841675868055554"); return g }(),
		},
		{
			"[ 1] 0.10 / 0.3 w 3 exp",
			func() Gyro { g, _ := NewFromString("0.10"); return g }(),
			func() Gyro { g, _ := NewFromString("0.3"); return g }(),
			3,
			func() Gyro { g, _ := NewFromString("0.333"); return g }(),
		},
		{
			"[ 2] 10 / 3 w 1 exp",
			func() Gyro { g, _ := NewFromString("10"); return g }(),
			func() Gyro { g, _ := NewFromString("3"); return g }(),
			1,
			func() Gyro { g, _ := NewFromString("3.3"); return g }(),
		},
		{
			"[ 3] 10 / 2 w 0 exp",
			func() Gyro { g, _ := NewFromString("10"); return g }(),
			func() Gyro { g, _ := NewFromString("2"); return g }(),
			0,
			func() Gyro { g, _ := NewFromString("5"); return g }(),
		},
		{
			"[ 4] 10 / -3 w 0 exp",
			func() Gyro { g, _ := NewFromString("10"); return g }(),
			func() Gyro { g, _ := NewFromString("-3"); return g }(),
			0,
			func() Gyro { g, _ := NewFromString("-3"); return g }(),
		},
		{
			"[ 5] 10 / 3 w 0 exp",
			func() Gyro { g, _ := NewFromString("10"); return g }(),
			func() Gyro { g, _ := NewFromString("3"); return g }(),
			0,
			func() Gyro { g, _ := NewFromString("3"); return g }(),
		},
		{
			"[ 6] -10 / -2 w 0 exp",
			func() Gyro { g, _ := NewFromString("-10"); return g }(),
			func() Gyro { g, _ := NewFromString("-2"); return g }(),
			0,
			func() Gyro { g, _ := NewFromString("5"); return g }(),
		},
		{
			"[ 7] 10 / -2 w 0 exp",
			func() Gyro { g, _ := NewFromString("-10"); return g }(),
			func() Gyro { g, _ := NewFromString("-2"); return g }(),
			0,
			func() Gyro { g, _ := NewFromString("5"); return g }(),
		},
		{
			"[ 8] -10 / 2 w 0 exp",
			func() Gyro { g, _ := NewFromString("-10"); return g }(),
			func() Gyro { g, _ := NewFromString("2"); return g }(),
			0,
			func() Gyro { g, _ := NewFromString("-5"); return g }(),
		},
		{
			"[ 9] 1 / 10 w 0 exp",
			func() Gyro { g, _ := NewFromString("-10"); return g }(),
			func() Gyro { g, _ := NewFromString("1"); return g }(),
			0,
			func() Gyro { g, _ := NewFromString("10"); return g }(),
		},
		{
			"[10] 0 / 10 w 0 exp",
			func() Gyro { g, _ := NewFromString("0"); return g }(),
			func() Gyro { g, _ := NewFromString("10"); return g }(),
			0,
			func() Gyro { g, _ := NewFromString("0"); return g }(),
		},
	}

	b.ResetTimer()

	for _, tc := range tests {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = tc.g1.DivRound(tc.g2, tc.precision)
			}
		})
	}
}
