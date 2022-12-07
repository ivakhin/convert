package convert

// SliceSafe convert []A to []B using function convert(A) B.
func SliceSafe[A, B any](in []A, convert func(A) B) []B {
	if in == nil {
		return nil
	}

	out := make([]B, 0, len(in))
	for _, a := range in {
		out = append(out, convert(a))
	}

	return out
}

// Slice convert []A to []B using function convert(A) (B, error).
func Slice[A, B any](in []A, convert func(A) (B, error)) ([]B, error) {
	if in == nil {
		return nil, nil
	}

	out := make([]B, 0, len(in))

	for _, a := range in {
		b, err := convert(a)
		if err != nil {
			return nil, err
		}

		out = append(out, b)
	}

	return out, nil
}

// SliceToMap convert []T to map[K]T using function key(T) K.
// If key() returns the same value for multiple slice elements, only the last element will be saved.
func SliceToMap[T any, K comparable](in []T, key func(T) K) map[K]T {
	out := make(map[K]T, len(in))
	for _, t := range in {
		out[key(t)] = t
	}

	return out
}

// MapToSlice make slice []T from map[K]T.
func MapToSlice[T any, K comparable](in map[K]T) []T {
	if in == nil {
		return nil
	}

	out := make([]T, 0, len(in))
	for _, t := range in {
		out = append(out, t)
	}

	return out
}

// SplitSlice convert []T to map[K][]T using function key(T) K.
func SplitSlice[T any, K comparable](in []T, key func(T) K) map[K][]T {
	out := make(map[K][]T, len(in))

	for _, t := range in {
		k := key(t)
		out[k] = append(out[k], t)
	}

	return out
}
