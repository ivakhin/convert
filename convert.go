package convert

// SliceSafe convert []A to []B using function convert(A) B.
func SliceSafe[A, B any](in []A, convert func(A) B) []B {
	if in == nil {
		return nil
	}

	out := make([]B, len(in))
	for i := range out {
		out[i] = convert(in[i])
	}

	return out
}

// Slice convert []A to []B using function convert(A) (B, error).
func Slice[A, B any](in []A, convert func(A) (B, error)) ([]B, error) {
	if in == nil {
		return nil, nil
	}

	out := make([]B, len(in))

	for i := range out {
		b, err := convert(in[i])
		if err != nil {
			return nil, err
		}

		out[i] = b
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

// SliceToMapWithConvert convert []T to map[K]S using function convert(T) (key K, value S).
// If convert() returns the same key for multiple slice elements, only the last element will be saved.
func SliceToMapWithConvert[T, S any, K comparable](in []T, convert func(T) (K, S)) map[K]S {
	out := make(map[K]S, len(in))

	for _, t := range in {
		key, value := convert(t)
		out[key] = value
	}

	return out
}

// MapToSlice make slice []T from map[K]T.
func MapToSlice[T any, K comparable](in map[K]T) []T {
	if in == nil {
		return nil
	}

	out := make([]T, len(in))
	i := 0

	for _, t := range in {
		out[i] = t
		i++
	}

	return out
}

// MapToSliceSafeWithConvert make slice []S from map[K]T, using function convert(T) S.
func MapToSliceSafeWithConvert[T, S any, K comparable](in map[K]T, convert func(T) S) []S {
	if in == nil {
		return nil
	}

	out := make([]S, len(in))
	i := 0

	for _, t := range in {
		out[i] = convert(t)
		i++
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

// SplitSliceWithConvert convert []T to map[K][]S using function key(T) (key K, value S).
func SplitSliceWithConvert[T, S any, K comparable](in []T, convert func(T) (K, S)) map[K][]S {
	out := make(map[K][]S, len(in))

	for _, t := range in {
		k, v := convert(t)
		out[k] = append(out[k], v)
	}

	return out
}

// Keys returns unordered slice of map keys.
func Keys[T any, K comparable](in map[K]T) []K {
	if in == nil {
		return nil
	}

	out := make([]K, 0, len(in))
	for key := range in {
		out = append(out, key)
	}

	return out
}
