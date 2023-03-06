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
BenchmarkSliceSafe/convert.SliceSafe-8         	                 6775	    170280 ns/op	  202721 B/op	    9901 allocs/op
BenchmarkSliceSafe/without_generics-8                            8214	    148134 ns/op	  202721 B/op	    9901 allocs/op

BenchmarkSlice/convert.Slice-8                 	                26539	     45114 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSlice/without_generics-8                               28632	     42019 ns/op	   81920 B/op	       1 allocs/op

BenchmarkSliceToMap/convert.SliceToMap-8       	                 2769	    431999 ns/op	  497690 B/op	    9902 allocs/op
BenchmarkSliceToMap/without_generics-8                           2788	    430209 ns/op	  497739 B/op	    9903 allocs/op

BenchmarkSliceToMapWithConvert/convert.SliceToMapWithConvert-8   2619	    461297 ns/op	  637042 B/op	    9903 allocs/op
BenchmarkSliceToMapWithConvert/without_generics-8                2652	    453764 ns/op	  637080 B/op	    9903 allocs/op

BenchmarkMapToSlice/convert.MapToSlice-8                        13512	     88878 ns/op	   81920 B/op	       1 allocs/op
BenchmarkMapToSlice/without_generics-8                          13496	     88630 ns/op	   81920 B/op	       1 allocs/op

BenchmarkSplitSlice/convert.SplitSlice-8                         2014	    596770 ns/op	  896667 B/op	   19903 allocs/op
BenchmarkSplitSlice/without_generics-8                           2050	    584057 ns/op	  896677 B/op	   19903 allocs/op
PASS
ok  	github.com/ivakhin/convert	14.925s
```