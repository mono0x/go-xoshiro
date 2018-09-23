package xoroshiro_test

import (
	"math/rand"
	"testing"

	"github.com/mono0x/go-xoshiro/splitmix64"
	"github.com/mono0x/go-xoshiro/xoroshiro128starstar"
	"github.com/mono0x/go-xoshiro/xoshiro256starstar"
	"github.com/seehuhn/mt19937"
)

func BenchmarkXoshiro256StarStar_NewSource(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = xoshiro256starstar.NewSource(1)
	}
}

func BenchmarkXoroshiro128StarStar_NewSource(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = xoroshiro128starstar.NewSource(1)
	}
}

func BenchmarkSplitMix64_NewSource(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = splitmix64.NewSource(1)
	}
}

func BenchmarkMathRand_NewSource(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rand.NewSource(1)
	}
}

func BenchmarkMT19937_New(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := mt19937.New()
		s.Seed(1)
	}
}
