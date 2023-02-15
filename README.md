![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)

#### Description
Go package with conversion methods based on generics.

```
SliceSafe  convert []A to   []B       using function convert(A) B
Slice      convert []A to   []B       using function convert(A) (B, error)
SliceToMap convert []T to   map[K]T   using function key(T) K
SplitSlice convert []T to   map[K][]T using function key(T) K
MapToSlice make    []T from map[K]T
```

#### Test coverage
```
github.com/ivakhin/convert/convert.go:4:	SliceSafe	100.0%
github.com/ivakhin/convert/convert.go:18:	Slice		100.0%
github.com/ivakhin/convert/convert.go:39:	SliceToMap	100.0%
github.com/ivakhin/convert/convert.go:49:	MapToSlice	100.0%
github.com/ivakhin/convert/convert.go:63:	SplitSlice	100.0%
total:						(statements)	100.0%
```

#### Benchmarks
```
~ go test -bench ./...
goos: darwin
goarch: arm64
pkg: github.com/ivakhin/convert
BenchmarkSliceSafe/with_generics_using_convert.SliceSafe-8              17679829                67.86 ns/op          160 B/op          1 allocs/op
BenchmarkSliceSafe/without_generics-8                                   18115087                65.48 ns/op          160 B/op          1 allocs/op

BenchmarkSlice/with_generics_using_convert.Slice-8                      25023067                47.66 ns/op           80 B/op          1 allocs/op
BenchmarkSlice/without_generics-8                                       25224568                46.96 ns/op           80 B/op          1 allocs/op

BenchmarkSliceToMap/with_generics_using_convert.SliceToMap-8             4460504               270.0 ns/op           468 B/op          2 allocs/op
BenchmarkSliceToMap/without_generics-8                                   4552183               264.2 ns/op           467 B/op          2 allocs/op

BenchmarkMapToSlice/with_generics_using_convert.MapToSlice-8            11272046               106.2 ns/op            80 B/op          1 allocs/op
BenchmarkMapToSlice/without_generics-8                                  11373109               105.3 ns/op            80 B/op          1 allocs/op

BenchmarkSplitSlice/with_generics_using_convert.SplitSlice-8             1830158               655.7 ns/op          1616 B/op         17 allocs/op
BenchmarkSplitSlice/without_generics-8                                   1844820               653.5 ns/op          1616 B/op         17 allocs/op

PASS
ok      github.com/ivakhin/convert      14.323s
```