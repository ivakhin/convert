package convert_test

import (
	"strconv"
	"testing"

	"github.com/ivakhin/convert"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSliceMust(t *testing.T) {
	t.Parallel()

	t.Run("nil", func(t *testing.T) {
		t.Parallel()

		var (
			in       []int
			expected []string
		)

		assert.Equal(t, convert.SliceSafe(in, strconv.Itoa), expected)
	})

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		var (
			in       = make([]int, 0)
			expected = make([]string, 0)
		)

		assert.Equal(t, convert.SliceSafe(in, strconv.Itoa), expected)
	})

	t.Run("regular", func(t *testing.T) {
		t.Parallel()

		var (
			in       = []int{1, 2, 3}
			expected = []string{"1", "2", "3"}
		)

		assert.Equal(t, convert.SliceSafe(in, strconv.Itoa), expected)
	})
}

func TestSlice(t *testing.T) {
	t.Parallel()

	t.Run("nil", func(t *testing.T) {
		t.Parallel()

		var (
			in       []string
			expected []int
		)

		actual, err := convert.Slice(in, strconv.Atoi)
		require.NoError(t, err)
		assert.Equal(t, actual, expected)
	})

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		var (
			in       = make([]string, 0)
			expected = make([]int, 0)
		)

		actual, err := convert.Slice(in, strconv.Atoi)
		require.NoError(t, err)
		assert.Equal(t, actual, expected)
	})

	t.Run("regular", func(t *testing.T) {
		t.Parallel()

		var (
			in       = []string{"1", "2", "3"}
			expected = []int{1, 2, 3}
		)

		actual, err := convert.Slice(in, strconv.Atoi)
		require.NoError(t, err)
		assert.Equal(t, actual, expected)
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		var (
			in       = []string{"a", "b", "c"}
			expected []int
		)

		actual, err := convert.Slice(in, strconv.Atoi)
		require.Error(t, err)
		assert.Equal(t, actual, expected)
	})
}

func TestSliceMap(t *testing.T) {
	t.Parallel()

	t.Run("nil", func(t *testing.T) {
		t.Parallel()

		var (
			in       map[string]int
			expected []int
		)

		actual := convert.MapToSlice(in)
		assert.Equal(t, actual, expected)
	})

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		var (
			in       = make(map[string]int)
			expected = make([]int, 0)
		)

		actual := convert.MapToSlice(in)
		assert.Equal(t, actual, expected)
	})

	t.Run("regular", func(t *testing.T) {
		t.Parallel()

		var (
			in       = map[string]int{"1": 1, "2": 2, "3": 3}
			expected = []int{1, 2, 3}
		)

		actual := convert.MapToSlice(in)
		assert.ElementsMatch(t, actual, expected)
	})
}

func TestMapSlice(t *testing.T) {
	t.Parallel()

	t.Run("nil", func(t *testing.T) {
		t.Parallel()

		var (
			in       []int
			expected = make(map[string]int)
		)

		actual := convert.SliceToMap(in, strconv.Itoa)
		assert.Equal(t, actual, expected)
	})

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		var (
			in       = make([]int, 0)
			expected = make(map[string]int)
		)

		actual := convert.SliceToMap(in, strconv.Itoa)
		assert.Equal(t, actual, expected)
	})

	t.Run("regular", func(t *testing.T) {
		t.Parallel()

		var (
			in       = []int{1, 2, 3}
			expected = map[string]int{"1": 1, "2": 2, "3": 3}
		)

		actual := convert.SliceToMap(in, strconv.Itoa)
		assert.Equal(t, actual, expected)
	})
}
