package i128

import (
	"math"
	"math/big"
)

const (
	maxUint64 = 1<<64 - 1
	MaxInt64  = 1<<63 - 1
	minInt64  = -1 << 63

	// WARNING: this can not be represented accurately as a float; attempting to
	// convert it to uint64 will overflow and cause weird truncation issues that
	// violate the principle of least astonishment.
	maxUint64Float  = float64(maxUint64)     // (1<<64) - 1
	wrapUint64Float = float64(maxUint64) + 1 // 1 << 64

	maxU128Float = float64(340282366920938463463374607431768211455)  // (1<<128) - 1
	maxI128Float = float64(170141183460469231731687303715884105727)  // (1<<127) - 1
	minI128Float = float64(-170141183460469231731687303715884105728) // -(1<<127)

	intSize = 32 << (^uint(0) >> 63)
)

var (
	MaxI128 = I128{hi: 0x7FFFFFFFFFFFFFFF, lo: 0xFFFFFFFFFFFFFFFF}
	MinI128 = I128{hi: 0x8000000000000000, lo: 0}
	MaxU128 = U128{hi: maxUint64, lo: maxUint64}

	zeroI128 I128
	zeroU128 U128

	big1 = new(big.Int).SetInt64(1)

	MaxBigU128, _ = new(big.Int).SetString("340282366920938463463374607431768211455", 10)
	MaxBigInt64   = new(big.Int).SetUint64(MaxInt64)
	MinBigInt64   = new(big.Int).SetInt64(minInt64)

	// I128 == MinI128.
	minI128AsU128 = U128{hi: 0x8000000000000000, lo: 0x0}

	maxRepresentableUint64Float = math.Nextafter(maxUint64Float, 0) // < (1<<64)

	maxRepresentableU128Float = math.Nextafter(float64(340282366920938463463374607431768211455), 0) // < (1<<128)
)
