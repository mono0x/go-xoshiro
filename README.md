# go-xoshiro

## Benchmark

### New

```
goos: darwin
goarch: amd64
pkg: github.com/mono0x/go-xoshiro
BenchmarkXoshiro256StarStar_NewSource-8         50000000                33.9 ns/op
BenchmarkXoroshiro128StarStar_NewSource-8       50000000                27.9 ns/op
BenchmarkSplitMix64_NewSource-8                 2000000000               0.30 ns/op
BenchmarkMathRand_NewSource-8                     200000             10572 ns/op
BenchmarkMT19937_New-8                           1000000              1078 ns/op
PASS
ok      github.com/mono0x/go-xoshiro    7.124s
```

### Int63

```
goos: darwin
goarch: amd64
pkg: github.com/mono0x/go-xoshiro
BenchmarkXoshiro256StarStar_Int63-8             200000000                7.63 ns/op
BenchmarkXoroshiro128StarStar_Int63-8           300000000                5.15 ns/op
BenchmarkSplitMix64_Int63-8                     2000000000               1.81 ns/op
BenchmarkMathRand_Int63-8                       300000000                4.63 ns/op
BenchmarkMT19937_Int63-8                        200000000                6.45 ns/op
PASS
ok      github.com/mono0x/go-xoshiro    11.996s
```

### Uint64

```
goos: darwin
goarch: amd64
pkg: github.com/mono0x/go-xoshiro
BenchmarkXoshiro256StarStar_Uint64-8            300000000                5.01 ns/op
BenchmarkXoroshiro128StarStar_Uint64-8          500000000                3.36 ns/op
BenchmarkSplitMix64_Uint64-8                    2000000000               1.80 ns/op
BenchmarkMathRand_Uint64-8                      300000000                4.54 ns/op
BenchmarkMT19937_Uint64-8                       200000000                6.20 ns/op
PASS
ok      github.com/mono0x/go-xoshiro    11.531s
```

