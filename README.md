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
BenchmarkSliceSafe/convert.SliceSafe-8                                              6806   169674 ns/op  202721 B/op   9901 allocs/op
BenchmarkSliceSafe/without_generics-8                                               7981   147823 ns/op  202721 B/op   9901 allocs/op

BenchmarkSlice/convert.Slice-8                                                     26721    44570 ns/op   81920 B/op      1 allocs/op
BenchmarkSlice/without_generics-8                                                  28833    41545 ns/op   81920 B/op      1 allocs/op

BenchmarkSliceToMap/convert.SliceToMap-8                                            2786   429046 ns/op  497685 B/op   9902 allocs/op
BenchmarkSliceToMap/without_generics-8                                              2820   426153 ns/op  497737 B/op   9903 allocs/op

BenchmarkSliceToMapWithConvert/convert.SliceToMapWithConvert-8                      2600   454602 ns/op  637053 B/op   9903 allocs/op
BenchmarkSliceToMapWithConvert/without_generics-8                                   2665   449447 ns/op  637045 B/op   9903 allocs/op

BenchmarkMapToSlice/convert.MapToSlice-8                                           13557    88551 ns/op   81920 B/op      1 allocs/op
BenchmarkMapToSlice/without_generics-8                                             13602    88389 ns/op   81920 B/op      1 allocs/op

BenchmarkMapToSliceSafeWithConvert/convert.MapToSliceSafeWithConvertSafe-8          4376   274839 ns/op  203441 B/op   9901 allocs/op
BenchmarkMapToSliceSafeWithConvert/without_generics-8                               4419   273862 ns/op  203441 B/op   9901 allocs/op

BenchmarkSplitSlice/convert.SplitSlice-8                                            2049   585630 ns/op  896683 B/op  19903 allocs/op
BenchmarkSplitSlice/without_generics-8                                              2062   577750 ns/op  896679 B/op  19903 allocs/op

BenchmarkSplitSliceWithConvert/convert.SplitSliceWithConvert-8                      1525   795688 ns/op  936384 B/op  19903 allocs/op
BenchmarkSplitSliceWithConvert/without_generics-8                                   1531   792817 ns/op  936346 B/op  19903 allocs/op
PASS
ok      github.com/ivakhin/convert 14.925s
```