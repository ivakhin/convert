![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)

#### Description
Go package with conversion methods based on generics.

```
SliceSafe  convert []A to   []B     without error using function convert(A) B
Slice      convert []A to   []B     using function convert(A) (B, error)
SliceToMap convert []T to   map[K]T using function key(T) K
MapToSlice make    []T from map[K]T
```

#### Test coverage
```
github.com/ivakhin/convert/convert.go:4:	SliceSafe	100.0%
github.com/ivakhin/convert/convert.go:18:	Slice		100.0%
github.com/ivakhin/convert/convert.go:38:	SliceToMap	100.0%
github.com/ivakhin/convert/convert.go:48:	MapToSlice	100.0%
total:						(statements)	100.0%
```

#### Benchmarks
```
~ go test -bench ./...
goos: darwin
goarch: arm64
pkg: github.com/ivakhin/convert

BenchmarkSliceSafe/with_generics_using_convert.SliceSafe-8              15142082                73.85 ns/op          160 B/op          1 allocs/op
BenchmarkSliceSafe/without_generics-8                                   17552926                68.28 ns/op          160 B/op          1 allocs/op

BenchmarkSlice/with_generics_using_convert.Slice-8                      23491008                50.93 ns/op           80 B/op          1 allocs/op
BenchmarkSlice/without_generics-8                                       25197400                47.38 ns/op           80 B/op          1 allocs/op

BenchmarkSliceToMap/with_generics_using_convert.SliceToMap-8             4386909               273.8 ns/op           467 B/op          2 allocs/op
BenchmarkSliceToMap/without_generics-8                                   4478426               267.4 ns/op           467 B/op          2 allocs/op

BenchmarkMapToSlice/with_generics_using_convert.MapToSlice-8            11184621               108.0 ns/op            80 B/op          1 allocs/op
BenchmarkMapToSlice/without_generics-8                                  11172945               106.0 ns/op            80 B/op          1 allocs/op
```