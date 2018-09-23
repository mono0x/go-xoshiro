package xoroshiro_test

import (
	"math/rand"
	"testing"

	"github.com/mono0x/go-xoshiro/splitmix64"
	"github.com/mono0x/go-xoshiro/xoroshiro128starstar"
	"github.com/mono0x/go-xoshiro/xoshiro256starstar"
	"github.com/seehuhn/mt19937"
)

func BenchmarkXoshiro256StarStar_Int63(b *testing.B) {
	s := xoshiro256starstar.NewSource(1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.Int63()
	}
}

func BenchmarkXoroshiro128StarStar_Int63(b *testing.B) {
	s := xoroshiro128starstar.NewSource(1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.Int63()
	}
}

func BenchmarkSplitMix64_Int63(b *testing.B) {
	s := splitmix64.NewSource(1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.Int63()
	}
}

func BenchmarkMathRand_Int63(b *testing.B) {
	s := rand.NewSource(1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.Int63()
	}
}

func BenchmarkMT19937_Int63(b *testing.B) {
	s := mt19937.New()
	s.Seed(1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.Int63()
	}
}
