package convert_test

import (
	"strconv"
	"testing"

	"github.com/ivakhin/convert"
)

func BenchmarkSliceSafe(b *testing.B) {
	in := intSlice()

	b.Run("convert.SliceSafe", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = convert.SliceSafe(in, strconv.Itoa)
		}
	})

	b.Run("without generics", func(b *testing.B) {
		b.ReportAllocs()

		convertFn := func(in []int) []string {
			out := make([]string, len(in))
			for i, v := range in {
				out[i] = strconv.Itoa(v)
			}

			return out
		}

		for i := 0; i < b.N; i++ {
			_ = convertFn(in)
		}
	})
}

func BenchmarkSlice(b *testing.B) {
	in := stringSlice()

	b.Run("convert.Slice", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_, _ = convert.Slice(in, strconv.Atoi)
		}
	})

	b.Run("without generics", func(b *testing.B) {
		b.ReportAllocs()

		convertFn := func(in []string) ([]int, error) {
			out := make([]int, len(in))
			for i, v := range in {
				value, err := strconv.Atoi(v)
				if err != nil {
					return nil, err //nolint:wrapcheck
				}

				out[i] = value
			}

			return out, nil
		}

		for i := 0; i < b.N; i++ {
			_, _ = convertFn(in)
		}
	})
}

func BenchmarkSliceToMap(b *testing.B) {
	in := intSlice()

	b.Run("convert.SliceToMap", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = convert.SliceToMap(in, strconv.Itoa)
		}
	})

	b.Run("without generics", func(b *testing.B) {
		b.ReportAllocs()

		convertFn := func(in []int) map[string]int {
			out := make(map[string]int, len(in))
			for _, v := range in {
				out[strconv.Itoa(v)] = v
			}

			return out
		}

		for i := 0; i < b.N; i++ {
			_ = convertFn(in)
		}
	})
}

func BenchmarkSliceToMapWithConvert(b *testing.B) {
	in := intSlice()

	b.Run("convert.SliceToMapWithConvert", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = convert.SliceToMapWithConvert(in, func(in int) (string, string) {
				out := strconv.Itoa(in)

				return out, out
			})
		}
	})

	b.Run("without generics", func(b *testing.B) {
		b.ReportAllocs()

		convertFn := func(in []int) map[string]string {
			out := make(map[string]string, len(in))
			for _, v := range in {
				str := strconv.Itoa(v)
				out[str] = str
			}

			return out
		}

		for i := 0; i < b.N; i++ {
			_ = convertFn(in)
		}
	})
}

func BenchmarkMapToSlice(b *testing.B) {
	in := convert.SliceToMap(intSlice(), strconv.Itoa)

	b.Run("convert.MapToSlice", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = convert.MapToSlice(in)
		}
	})

	b.Run("without generics", func(b *testing.B) {
		b.ReportAllocs()

		convertFn := func(in map[string]int) []int {
			out := make([]int, 0, len(in))
			for _, v := range in {
				out = append(out, v)
			}

			return out
		}

		for i := 0; i < b.N; i++ {
			_ = convertFn(in)
		}
	})
}

func BenchmarkMapToSliceSafeWithConvert(b *testing.B) {
	in := convert.SliceToMap(intSlice(), func(i int) int {
		return i
	})

	b.Run("convert.MapToSliceSafeWithConvertSafe", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = convert.MapToSliceSafeWithConvert(in, strconv.Itoa)
		}
	})

	b.Run("without generics", func(b *testing.B) {
		b.ReportAllocs()

		convertFn := func(in map[int]int) []string {
			out := make([]string, 0, len(in))
			for _, v := range in {
				out = append(out, strconv.Itoa(v))
			}

			return out
		}

		for i := 0; i < b.N; i++ {
			_ = convertFn(in)
		}
	})
}

func BenchmarkSplitSlice(b *testing.B) {
	in := intSlice()

	b.Run("convert.SplitSlice", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = convert.SplitSlice(in, strconv.Itoa)
		}
	})

	b.Run("without generics", func(b *testing.B) {
		b.ReportAllocs()

		convertFn := func(in []int) map[string][]int {
			out := make(map[string][]int, len(in))
			for _, v := range in {
				key := strconv.Itoa(v)
				out[key] = append(out[key], v)
			}

			return out
		}

		for i := 0; i < b.N; i++ {
			_ = convertFn(in)
		}
	})
}

func intSlice() []int {
	in := make([]int, 10000)
	for i := range in {
		in[i] = i
	}

	return in
}

func stringSlice() []string {
	in := make([]string, 10000)
	for i := range in {
		in[i] = strconv.Itoa(i)
	}

	return in
}
