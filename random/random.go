package common

import (
	"math/rand"
	"time"
)

var source = rand.NewSource(time.Now().UnixNano())
var random = rand.New(source)

func RandomFloat64(from, to float64) float64 {
	if to-from <= 0 {
		return to
	}
	return random.Float64()*(to-from) + from
}

func RandomInt64(from, to int64) int64 {
	if to-from <= 0 {
		return to
	}
	return random.Int63n(to-from+1) + from
}

func RandomInt(from, to int) int {
	if to-from <= 0 {
		return to
	}
	return random.Intn(to-from+1) + from
}
