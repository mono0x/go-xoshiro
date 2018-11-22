package xoshiro256starstar

import (
	"math/bits"

	"github.com/mono0x/go-xoshiro/splitmix64"
)

// http://xoshiro.di.unimi.it/xoshiro256starstar.c

type Source struct {
	state [4]uint64
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
	s.state[2] = splitMix.Uint64()
	s.state[3] = splitMix.Uint64()
}

func (s *Source) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s *Source) Uint64() uint64 {
	result := bits.RotateLeft64(s.state[1]*5, 7) * 9

	t := s.state[1] << 17

	s.state[2] ^= s.state[0]
	s.state[3] ^= s.state[1]
	s.state[1] ^= s.state[2]
	s.state[0] ^= s.state[3]

	s.state[2] ^= t

	s.state[3] = bits.RotateLeft64(s.state[3], 45)

	return result
}

var jump = [4]uint64{0x180ec6d33cfd0aba, 0xd5a61266f0c9392c, 0xa9582618e03fc9aa, 0x39abdc4529b1661c}

// This is the jump function for the generator. It is equivalent
// to 2^128 calls to next(); it can be used to generate 2^128
// non-overlapping subsequences for parallel computations.
func (s *Source) Jump() {
	var s0, s1, s2, s3 uint64
	for _, j := range jump {
		for b := uint64(0); b < 64; b++ {
			if (j & (1 << b)) != 0 {
				s0 ^= s.state[0]
				s1 ^= s.state[1]
				s2 ^= s.state[2]
				s3 ^= s.state[3]
			}
			s.Uint64()
		}
	}

	s.state[0] = s0
	s.state[1] = s1
	s.state[2] = s2
	s.state[3] = s3
}

var longJump = [4]uint64{0x76e15d3efefdcbbf, 0xc5004e441c522fb3, 0x77710069854ee241, 0x39109bb02acbe635}

// This is the long-jump function for the generator. It is equivalent to
// 2^192 calls to next(); it can be used to generate 2^64 starting points,
// from each of which jump() will generate 2^64 non-overlapping
// subsequences for parallel distributed computations.
func (s *Source) LongJump() {
	var s0, s1, s2, s3 uint64
	for _, j := range longJump {
		for b := uint64(0); b < 64; b++ {
			if (j & (1 << b)) != 0 {
				s0 ^= s.state[0]
				s1 ^= s.state[1]
				s2 ^= s.state[2]
				s3 ^= s.state[3]
			}
			s.Uint64()
		}
	}

	s.state[0] = s0
	s.state[1] = s1
	s.state[2] = s2
	s.state[3] = s3
}
