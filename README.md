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
BenchmarkSliceSafe/with_generics_using_convert.SliceSafe-8         	    7714	    150328 ns/op	  202721 B/op	    9901 allocs/op
BenchmarkSliceSafe/without_generics-8                              	    8383	    143837 ns/op	  202721 B/op	    9901 allocs/op

BenchmarkSlice/with_generics_using_convert.Slice-8                 	   25006	     47835 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSlice/without_generics-8                                  	   26739	     44721 ns/op	   81920 B/op	       1 allocs/op

BenchmarkSliceToMap/with_generics_using_convert.SliceToMap-8       	    2748	    435984 ns/op	  497743 B/op	    9903 allocs/op
BenchmarkSliceToMap/without_generics-8                             	    2766	    436712 ns/op	  497745 B/op	    9903 allocs/op

BenchmarkMapToSlice/with_generics_using_convert.MapToSlice-8       	   13453	     88736 ns/op	   81920 B/op	       1 allocs/op
BenchmarkMapToSlice/without_generics-8                             	   13506	     88905 ns/op	   81920 B/op	       1 allocs/op

BenchmarkSplitSlice/with_generics_using_convert.SplitSlice-8       	    1995	    592693 ns/op	  896686 B/op	   19903 allocs/op
BenchmarkSplitSlice/without_generics-8                             	    2028	    588764 ns/op	  896702 B/op	   19903 allocs/op
PASS
ok  	github.com/ivakhin/convert	14.925s
```