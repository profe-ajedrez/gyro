package gyro

import (
	"fmt"
	"testing"

	"github.com/profe-ajedrez/gyro/i128"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		name     string
		coeff    i128.I128
		exp      int32
		expected Gyro
	}{
		{
			name:     "Zero coeff and exp",
			coeff:    i128.I128FromRaw(0, 123),
			exp:      -2,
			expected: Gyro{coeff: func() i128.I128 { i := i128.I128FromInt(123); return i }(), exp: -2},
		},
		{
			name:     "Zero coeff and exp",
			coeff:    i128.I128FromRaw(0, 456),
			exp:      -4,
			expected: Gyro{coeff: func() i128.I128 { i := i128.I128FromInt(456); return i }(), exp: -4},
		},
		{
			name:     "Zero coeff and exp",
			coeff:    i128.I128{},
			exp:      0,
			expected: Gyro{coeff: func() i128.I128 { i := i128.I128FromInt(0); return i }(), exp: 0},
		},
		{
			name:     "Positive coeff and exp",
			coeff:    i128.I128FromInt(123456789),
			exp:      5,
			expected: Gyro{coeff: func() i128.I128 { i := i128.MustI128FromString("12345678900000"); return i }(), exp: 0},
		},
		{
			name:     "Negative coeff and exp",
			coeff:    i128.I128FromInt(-987654321),
			exp:      -5,
			expected: Gyro{coeff: func() i128.I128 { i := i128.MustI128FromString("-987654321"); return i }(), exp: -5},
		},
	}

	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := New(tc.coeff, tc.exp)
			fmt.Println(actual.String())
			if !tc.expected.coeff.Equal(actual.coeff) || actual.exp != tc.expected.exp {
				t.Errorf("[test case %d] New(%v, %d) = [%v][%v], expected [%v][%v]", i, tc.coeff, tc.exp, actual.coeff, actual.exp, tc.expected.coeff, tc.expected.exp)
			}
		})
	}
}

func TestNewFromString(t *testing.T) {
	testCases := []struct {
		name       string
		input      string
		shouldfail bool
		expected   Gyro
		err        error
	}{
		{
			name:       "Valid decimal with negative exponent with large fraction",
			input:      "-782378237823782378.12356784567890754433556677999999999999999",
			shouldfail: false,
			expected:   Gyro{coeff: i128.MustI128FromString("-782378237823782378123567845678907"), exp: -16},
			err:        nil,
		},
		{
			name:       "Valid decimal with negative exponent with large fraction",
			input:      "-782378237823782378.12356784567890754433556677999999999999999",
			shouldfail: false,
			expected:   Gyro{coeff: i128.MustI128FromString("-782378237823782378123567845678907"), exp: -16},
			err:        nil,
		},
		{
			name:       "Valid decimal with negative exponent with large fraction",
			input:      "782378237823782378.12356784567890754433556677999999999999999",
			shouldfail: false,
			expected:   Gyro{coeff: i128.MustI128FromString("782378237823782378123567845678907"), exp: -16},
			err:        nil,
		},
		{
			name:       "Valid negative integer",
			input:      "-67890",
			shouldfail: false,
			expected:   Gyro{coeff: i128.I128FromInt(-67890), exp: 0},
			err:        nil,
		},
		{
			name:       "Empty string",
			input:      "",
			shouldfail: true,
			expected:   Gyro{},
			err:        NewGyroInvalidErr("couldnt convert an empty space in a decimal value"),
		},
		{
			name:       "Valid integer",
			input:      "12345",
			shouldfail: false,
			expected:   Gyro{coeff: i128.I128FromInt(12345), exp: 0},
			err:        nil,
		},
		{
			name:       "Valid decimal with negative exponent",
			input:      "12.345",
			shouldfail: false,
			expected:   Gyro{coeff: i128.MustI128FromString("12345"), exp: -3},
			err:        nil,
		},
		{
			name:       "Valid decimal with negative exponent with large fraction",
			input:      "67890.12356784567890754433556677",
			shouldfail: false,
			expected:   Gyro{coeff: i128.MustI128FromString("67890123567845678907"), exp: -16},
			err:        nil,
		},
		{
			name:       "Multiple decimal points",
			input:      "12.34.56",
			shouldfail: true,
			expected:   Gyro{},
			err:        NewGyroInvalidErr("only decimal strings are supported. 12.34.56 has more tha 1 decimal point"),
		},
		{
			name:       "Invalid character",
			input:      "12a34",
			shouldfail: true,
			expected:   Gyro{},
			err:        NewGyroInvalidErr("invalid character in string: 12a34"),
		},
	}

	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := NewFromString(tc.input)

			if !tc.shouldfail && err != nil && tc.err == nil {
				t.Errorf("[test case %d] NewFromString(%q) returned unexpected error: %v", i, tc.input, err)
			} else if tc.shouldfail && err == nil && tc.err != nil {
				t.Errorf("[test case %d] NewFromString(%q) expected error %v, but got nil", i, tc.input, tc.err)
			} else if !tc.shouldfail && err != nil && err.Error() != tc.err.Error() {
				t.Errorf("[test case %d] NewFromString(%q) expected error %v, but got %v", i, tc.input, tc.err, err)
			} else if !tc.shouldfail && !actual.coeff.Equal(tc.expected.coeff) || actual.exp != tc.expected.exp {
				// fmt.Println(actual.coeff.String())
				// fmt.Println(actual.exp)
				// fmt.Println(tc.expected.coeff.String())
				// fmt.Println(tc.expected.exp)
				t.Errorf("[test case %d] NewFromString(%q) = %v, expected %v", i, tc.input, actual, tc.expected)
			}
		})
	}
}

func TestGyroString(t *testing.T) {
	testCases := []struct {
		name     string
		coeff    i128.I128
		exp      int32
		expected string
	}{
		{
			name:     "Zero coeff and exp",
			coeff:    i128.I128{},
			exp:      0,
			expected: "0",
		},
		{
			name:     "Positive coeff and zero exp",
			coeff:    i128.I128FromInt(12345),
			exp:      0,
			expected: "12345",
		},
		{
			name:     "Negative coeff and zero exp",
			coeff:    i128.I128FromInt(-67890),
			exp:      0,
			expected: "-67890",
		},
		{
			name:     "Positive coeff and positive exp",
			coeff:    i128.I128FromInt(12345),
			exp:      3,
			expected: "12345000",
		},
		{
			name:     "Negative coeff and positive exp",
			coeff:    i128.I128FromInt(-67890),
			exp:      5,
			expected: "-6789000000",
		},
		{
			name:     "Positive coeff and negative exp",
			coeff:    i128.I128FromInt(12345),
			exp:      -2,
			expected: "123.45",
		},
		{
			name:     "Negative coeff and negative exp",
			coeff:    i128.I128FromInt(-67890),
			exp:      -4,
			expected: "-6.7890",
		},
	}

	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gyro := New(tc.coeff, tc.exp)
			actual := gyro.String()
			if actual != tc.expected {
				t.Errorf("[test case %d] Gyro.String() = %s, expected %s", i, actual, tc.expected)
			}
		})
	}
}

func TestGyroFloat64(t *testing.T) {
	tests := []struct {
		name     string
		g        Gyro
		expected float64
	}{
		{
			name:     "Positive exponent",
			g:        Gyro{coeff: i128.I128FromU64(12345), exp: -2},
			expected: 123.45,
		},
		{
			name:     "Negative exponent",
			g:        Gyro{coeff: i128.I128FromU64(123451), exp: -1},
			expected: 12345.1,
		},
		{
			name:     "Zero exponent",
			g:        Gyro{coeff: i128.I128FromU64(12345), exp: 0},
			expected: 12345.0,
		},
		{
			name:     "Large positive exponent",
			g:        Gyro{coeff: i128.I128FromU64(12345), exp: -20},
			expected: 1.2345e-16,
		},
		{
			name:     "Large negative exponent",
			g:        Gyro{coeff: i128.I128FromU64(12345), exp: 20},
			expected: 1.2345e+24,
		},
	}

	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.g.Float64()
			if result != test.expected {
				t.Errorf("[test case %d] Float64() = %v, expected %v  <%f>", i, result, test.expected, result)
			}

			fmt.Printf("%f   %f\n", result, test.expected)
		})
	}
}

func TestGyroAdd(t *testing.T) {
	testCases := []struct {
		name     string
		g1       Gyro
		g2       Gyro
		expected Gyro
	}{
		{
			name:     "Different exponents",
			g1:       NewFromInt64Raw(123, 2),
			g2:       NewFromInt64Raw(456, 4),
			expected: NewFromInt64Raw(4572300, 0),
		},
		{
			name:     "Positive and zero",
			g1:       NewFromInt64Raw(123, 0),
			g2:       New(i128.I128{}, 0),
			expected: NewFromInt64Raw(123, 0),
		},
		{
			name:     "Zero and zero",
			g1:       New(i128.I128{}, 0),
			g2:       New(i128.I128{}, 0),
			expected: New(i128.I128{}, 0),
		},
		{
			name:     "Negative and zero",
			g1:       NewFromInt64Raw(-456, 0),
			g2:       New(i128.I128{}, 0),
			expected: NewFromInt64Raw(-456, 0),
		},
		{
			name:     "Positive and positive",
			g1:       NewFromInt64Raw(123, 2),
			g2:       NewFromInt64Raw(456, 2),
			expected: NewFromInt64Raw(579, 2),
		},
		{
			name:     "Negative and negative",
			g1:       NewFromInt64Raw(-123, 2),
			g2:       NewFromInt64Raw(-456, 2),
			expected: NewFromInt64Raw(-579, 2),
		},
		{
			name:     "Positive and negative",
			g1:       NewFromInt64Raw(123, 2),
			g2:       NewFromInt64Raw(-456, 2),
			expected: NewFromInt64Raw(-333, 2),
		},
		{
			name:     "Negative and positive",
			g1:       NewFromInt64Raw(-123, 2),
			g2:       NewFromInt64Raw(456, 2),
			expected: NewFromInt64Raw(333, 2),
		},
		{
			name:     "Large values",
			g1:       func() Gyro { g, _ := NewFromString("123456789012345678901"); return g }(),
			g2:       func() Gyro { g, _ := NewFromString("987654321098765432109"); return g }(),
			expected: func() Gyro { g, _ := NewFromString("1111111110111111111010"); return g }(),
		},
	}

	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.g1.Add(tc.g2)
			if !result.Equal(tc.expected) {
				t.Errorf("[test case %d]Add(%v, %v) = %v, expected %v", i, tc.g1, tc.g2, result, tc.expected)
			}
		})
	}
}

func TestGyroSub(t *testing.T) {
	testCases := []struct {
		name     string
		g1       Gyro
		g2       Gyro
		expected Gyro
	}{
		{
			name:     "Different exponents",
			g1:       NewFromInt64Raw(12300, -2),
			g2:       NewFromInt64Raw(4560000, -4),
			expected: NewFromInt64Raw(-333, 0),
		},
		{
			name:     "Zero minus zero",
			g1:       New(i128.I128{}, 0),
			g2:       New(i128.I128{}, 0),
			expected: New(i128.I128{}, 0),
		},
		{
			name:     "Positive minus zero",
			g1:       NewFromInt64Raw(123, 0),
			g2:       New(i128.I128{}, 0),
			expected: NewFromInt64Raw(123, 0),
		},
		{
			name:     "Negative minus zero",
			g1:       NewFromInt64Raw(-456, 0),
			g2:       New(i128.I128{}, 0),
			expected: NewFromInt64Raw(-456, 0),
		},
		{
			name:     "Positive minus positive (larger - smaller)",
			g1:       NewFromInt64Raw(456, 2),
			g2:       NewFromInt64Raw(123, 2),
			expected: NewFromInt64Raw(333, 2),
		},
		{
			name:     "Positive minus positive (smaller - larger)",
			g1:       NewFromInt64Raw(123, 2),
			g2:       NewFromInt64Raw(456, 2),
			expected: NewFromInt64Raw(-333, 2),
		},
		{
			name:     "Negative minus negative (larger - smaller)",
			g1:       NewFromInt64Raw(-123, 2),
			g2:       NewFromInt64Raw(-456, 2),
			expected: NewFromInt64Raw(333, 2),
		},
		{
			name:     "Negative minus negative (smaller - larger)",
			g1:       NewFromInt64Raw(-456, 2),
			g2:       NewFromInt64Raw(-123, 2),
			expected: NewFromInt64Raw(-333, 2),
		},
		{
			name:     "Positive minus negative",
			g1:       NewFromInt64Raw(123, 2),
			g2:       NewFromInt64Raw(-456, 2),
			expected: NewFromInt64Raw(579, 2),
		},
		{
			name:     "Negative minus positive",
			g1:       NewFromInt64Raw(-123, 2),
			g2:       NewFromInt64Raw(456, 2),
			expected: NewFromInt64Raw(-579, 2),
		},
		{
			name:     "Large values",
			g1:       func() Gyro { g, _ := NewFromString("123456789012345678901"); return g }(),
			g2:       func() Gyro { g, _ := NewFromString("987654321098765432109"); return g }(),
			expected: func() Gyro { g, _ := NewFromString("-864197532086419753208"); return g }(),
		},
	}

	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.g1.Sub(tc.g2)
			if !result.Equal(tc.expected) {
				t.Errorf("[test case %d] Sub(%v, %v) = %v, expected %v", i, tc.g1, tc.g2, result, tc.expected)
			}
		})
	}
}

func TestGyroMul(t *testing.T) {
	testCases := []struct {
		name     string
		g1       Gyro
		g2       Gyro
		expected Gyro
	}{
		{
			name:     "decimal values. different exponents",
			g1:       NewFromInt64Raw(-123, -2),
			g2:       NewFromInt64Raw(-456, -4),
			expected: NewFromInt64Raw(56088, -6),
		},
		{
			name:     "same exponents.",
			g1:       NewFromInt64Raw(-123, 2),
			g2:       NewFromInt64Raw(-456, 2),
			expected: NewFromInt64Raw(560880000, 0),
		},
		{
			name:     "Different exponents",
			g1:       NewFromInt64Raw(123, 2),
			g2:       NewFromInt64Raw(456, 4),
			expected: NewFromInt64Raw(56088000000, 0),
		},
		{
			name:     "Different exponents. first factor negative",
			g1:       NewFromInt64Raw(-123, 2),
			g2:       NewFromInt64Raw(456, 4),
			expected: NewFromInt64Raw(-56088000000, 0),
		},
		{
			name:     "Different exponents. second factor negative",
			g1:       NewFromInt64Raw(123, 2),
			g2:       NewFromInt64Raw(-456, 4),
			expected: NewFromInt64Raw(-56088000000, 0),
		},
		{
			name:     "Different exponents. both factor negative",
			g1:       NewFromInt64Raw(-123, 2),
			g2:       NewFromInt64Raw(-456, 4),
			expected: NewFromInt64Raw(56088000000, 0),
		},
	}

	for i, tc := range testCases {
		result := tc.g1.Mul(tc.g2)

		if !result.Equal(tc.expected) {
			t.Errorf("[test case %d] Mul(%v, %v) = <%v %v>, expected <%v, %v>", i, tc.g1, tc.g2, result.coeff, result.exp, tc.expected.coeff, tc.expected.exp)
			t.FailNow()
		}
	}

}

// Suponiendo que tienes una implementación de NewFromString que crea un Gyro desde un string
func TestDivRound(t *testing.T) {

	tests := []struct {
		g1        string
		g2        string
		precision int32
		expected  string
	}{
		{"-10.10112212", "2.304", 16, "-4.3841675868055554"},
		{"0.10", "0.3", 3, "0.333"},
		{"10", "3", 1, "3.3"},
		{"10", "2", 0, "5"},
		{"10", "-3", 0, "-3"},
		{"10", "3", 0, "3"},
		{"-10", "-2", 0, "5"},
		{"10", "-2", 0, "-5"},
		{"-10", "2", 0, "-5"},
		{"10", "1", 0, "10"},
		{"0", "10", 0, "0"},
	}

	for i, test := range tests {
		g1, err := NewFromString(test.g1)

		if err != nil {
			t.Errorf("[test case %d] Error creating Gyro from string: %s", i, test.g1)
		}

		g2, err := NewFromString(test.g2)

		if err != nil {
			t.Errorf("[test case %d] Error creating Gyro from string: %s", i, test.g2)
		}

		result := g1.DivRound(g2, test.precision)

		expected, err := NewFromString(test.expected)
		if err != nil {
			t.Errorf("[test case %d] Error creating expected Gyro from string: %s", i, test.expected)
		}

		if !result.Equal(expected) {
			t.Errorf("[test case %d] For %s / %s with precision %d, expected %s but got %s",
				i, test.g1, test.g2, test.precision, test.expected, result)
			t.FailNow()
		}
	}
}

func TestDivRoundByZero(t *testing.T) {
	g1, err := NewFromString("10")

	if err != nil {
		t.Errorf("Error creating Gyro from string: %s", "10")
	}

	g2, err := NewFromString("0")

	if err != nil {
		t.Errorf("Error creating Gyro from string: %s", "0")
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when dividing by zero")
		}
	}()

	g1.DivRound(g2, 0)
}

func TestGyro_Round(t *testing.T) {
	tests := []struct {
		name string
		num  func() Gyro
		want Gyro
	}{
		{
			name: "<4 decimals up>",
			num:  func() Gyro { g, _ := NewFromString("134.17356"); return g.Round(4) },
			want: func() Gyro { g, _ := NewFromString("134.1736"); return g }(),
		},
		{
			name: "<3 decimals up>",
			num:  func() Gyro { g, _ := NewFromString("134.17356"); return g.Round(3) },
			want: func() Gyro { g, _ := NewFromString("134.174"); return g }(),
		},
		{
			name: "<2 decimals up>",
			num:  func() Gyro { g, _ := NewFromString("1.178"); return g.Round(2) },
			want: func() Gyro { g, _ := NewFromString("1.18"); return g }(),
		},
		{
			name: "<2 decimals down>",
			num:  func() Gyro { g, _ := NewFromString("1.173"); return g.Round(2) },
			want: func() Gyro { g, _ := NewFromString("1.17"); return g }(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.num()
			if !got.Equal(tt.want) {
				t.Errorf("Gyro.Round() = %v, want %v", got, tt.want)
			}
		})
	}
}
