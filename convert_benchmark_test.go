package convert_test

import (
	"strconv"
	"testing"

	"github.com/ivakhin/convert"
)

func BenchmarkSliceSafe(b *testing.B) {
	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	b.Run("with generics using convert.SliceSafe", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = convert.SliceSafe(in, strconv.Itoa)
		}
	})

	b.Run("without generics", func(b *testing.B) {
		b.ReportAllocs()

		convertFn := func(in []int) []string {
			out := make([]string, 0, len(in))
			for _, i := range in {
				out = append(out, strconv.Itoa(i))
			}

			return out
		}

		for i := 0; i < b.N; i++ {
			_ = convertFn(in)
		}
	})
}

func BenchmarkSlice(b *testing.B) {
	in := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

	b.Run("with generics using convert.Slice", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_, _ = convert.Slice(in, strconv.Atoi)
		}
	})

	b.Run("without generics", func(b *testing.B) {
		b.ReportAllocs()

		convertFn := func(in []string) ([]int, error) {
			out := make([]int, 0, len(in))
			for _, i := range in {
				value, err := strconv.Atoi(i)
				if err != nil {
					return nil, err //nolint:wrapcheck
				}

				out = append(out, value)
			}

			return out, nil
		}

		for i := 0; i < b.N; i++ {
			_, _ = convertFn(in)
		}
	})
}

func BenchmarkSliceToMap(b *testing.B) {
	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	b.Run("with generics using convert.SliceToMap", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = convert.SliceToMap(in, strconv.Itoa)
		}
	})

	b.Run("without generics", func(b *testing.B) {
		b.ReportAllocs()

		convertFn := func(in []int) map[string]int {
			out := make(map[string]int, len(in))
			for _, i := range in {
				out[strconv.Itoa(i)] = i
			}

			return out
		}

		for i := 0; i < b.N; i++ {
			_ = convertFn(in)
		}
	})
}

func BenchmarkMapToSlice(b *testing.B) {
	in := map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "0": 0}

	b.Run("with generics using convert.MapToSlice", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = convert.MapToSlice(in)
		}
	})

	b.Run("without generics", func(b *testing.B) {
		b.ReportAllocs()

		convertFn := func(in map[string]int) []int {
			out := make([]int, 0, len(in))
			for _, i := range in {
				out = append(out, i)
			}

			return out
		}

		for i := 0; i < b.N; i++ {
			_ = convertFn(in)
		}
	})
}

func BenchmarkSplitSlice(b *testing.B) {
	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5}

	b.Run("with generics using convert.SplitSlice", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = convert.SplitSlice(in, strconv.Itoa)
		}
	})

	b.Run("without generics", func(b *testing.B) {
		b.ReportAllocs()

		convertFn := func(in []int) map[string][]int {
			out := make(map[string][]int, len(in))
			for _, i := range in {
				key := strconv.Itoa(i)
				out[key] = append(out[key], i)
			}

			return out
		}

		for i := 0; i < b.N; i++ {
			_ = convertFn(in)
		}
	})
}
