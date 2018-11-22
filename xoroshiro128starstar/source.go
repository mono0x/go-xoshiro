package xoroshiro128starstar

import (
	"math/bits"

	"github.com/mono0x/go-xoshiro/splitmix64"
)

// http://xoshiro.di.unimi.it/xoroshiro128starstar.c

type Source struct {
	state [2]uint64
}

func NewSource(seed int64) *Source {
	s := &Source{}
	s.Seed(seed)
	return s
}

func (s *Source) Seed(seed int64) {
	splitMix := splitmix64.NewSource(seed)
	s.state[0] = splitMix.Uint64()
	s.state[1] = splitMix.Uint64()
}

func (s *Source) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s *Source) Uint64() uint64 {
	s0 := s.state[0]
	s1 := s.state[1]

	result := bits.RotateLeft64(s0*5, 7) * 9

	s1 ^= s0
	s.state[0] = bits.RotateLeft64(s0, 24) ^ s1 ^ (s1 << 16)
	s.state[1] = bits.RotateLeft64(s1, 37)

	return result
}

var jump = [4]uint64{0xdf900294d8f554a5, 0x170865df4b3201fc}

// This is the jump function for the generator. It is equivalent
// to 2^64 calls to next(); it can be used to generate 2^64
// non-overlapping subsequences for parallel computations.
func (s *Source) Jump() {
	var s0, s1 uint64
	for _, j := range jump {
		for b := uint64(0); b < 64; b++ {
			if (j & (1 << b)) != 0 {
				s0 ^= s.state[0]
				s1 ^= s.state[1]
			}
			s.Uint64()
		}
	}

	s.state[0] = s0
	s.state[1] = s1
}

var longJump = [4]uint64{0xd2a98b26625eee7b, 0xdddf9b1090aa7ac1}

// This is the long-jump function for the generator. It is equivalent to
// 2^96 calls to next(); it can be used to generate 2^32 starting points,
// from each of which jump() will generate 2^32 non-overlapping
// subsequences for parallel distributed computations.
func (s *Source) LongJump() {
	var s0, s1 uint64
	for _, j := range longJump {
		for b := uint64(0); b < 64; b++ {
			if (j & (1 << b)) != 0 {
				s0 ^= s.state[0]
				s1 ^= s.state[1]
			}
			s.Uint64()
		}
	}

	s.state[0] = s0
	s.state[1] = s1
}
