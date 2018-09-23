package splitmix64

// http://xoshiro.di.unimi.it/splitmix64.c

type Source struct {
	x uint64
}

func NewSource(seed int64) *Source {
	s := &Source{}
	s.Seed(seed)
	return s
}

func (s *Source) Seed(seed int64) {
	s.x = uint64(seed)
}

func (s *Source) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1 << 63))
}

func (s *Source) Uint64() uint64 {
	s.x += uint64(0x9E3779B97F4A7C15)
	z := s.x
	z = (z ^ (z >> 30)) * uint64(0xBF58476D1CE4E5B9)
	z = (z ^ (z >> 27)) * uint64(0x94D049BB133111EB)
	return z ^ (z >> 31)
}
