package generics

import (
	"cmp"
	"encoding/json"
)

// MapKeysAndValues returns the keys and values of a map as separate slices
func MapKeysAndValues[K comparable, V any](m map[K]V) (keys []K, values []V) {
	length, i := len(m), 0
	keys = make([]K, length)
	values = make([]V, length)

	for k, v := range m {
		keys[i] = k
		values[i] = v
		i++
	}

	return
}

// MapKeys returns the keys from m as slice
func MapKeys[K comparable, V any](m map[K]V) (keys []K) {
	for k := range m {
		keys = append(keys, k)
	}
	return
}

// MapValues returns the values from m as slice
func MapValues[K comparable, V any](m map[K]V) (values []V) {
	for _, v := range m {
		values = append(values, v)
	}
	return
}

// SliceToMap transforms a slice of comparables to a custom map
func SliceToMap[K comparable, M ~[]K, T any](s M, defaultValue T) (res map[K]T) {
	res = make(map[K]T)
	for i := range s {
		res[s[i]] = defaultValue
	}
	return
}

// SliceToBoolMap transforms a slice of comparables to a boolean map with all elements set to 'true'
func SliceToBoolMap[K comparable, M ~[]K](s M) (res map[K]bool) {
	return SliceToMap(s, true)
}

// ConvertByJSON is a shortcut to the common use of the functions Marshal and Unmarshal from the json package
func ConvertByJSON[T any](src any) (dest T, err error) {
	var data []byte
	if data, err = json.Marshal(src); err != nil {
		return
	}
	err = json.Unmarshal(data, &dest)
	return
}

// Clamp limits the value of val, if it's lower than low, it returns low,
// if val is higher that high, it returns high, otherwise returns val
func Clamp[T cmp.Ordered](low, high, val T) T {
	if val < low {
		return low
	}
	if val > high {
		return high
	}
	return val
}

// InBetween checks if val belongs to (low, high)
func InBetween[T cmp.Ordered](low, high, val T) bool {
	return low < val && high > val
}

// InBetweenIncluding checks if val belongs to [low, high]
func InBetweenIncluding[T cmp.Ordered](low, high, val T) bool {
	return low <= val && high >= val
}

// InBetweenIncludingLow checks if val belongs to [low, high)
func InBetweenIncludingLow[T cmp.Ordered](low, high, val T) bool {
	return low <= val && high > val
}

// InBetweenIncludingHigh checks if val belongs to (low, high]
func InBetweenIncludingHigh[T cmp.Ordered](low, high, val T) bool {
	return low < val && high >= val
}

// Min returns the minimum element of values
func Min[T cmp.Ordered](values ...T) T {
	return MinByFunc(func(a, b T) bool { return a < b }, values...)
}

// MinByFunc returns the minimum element of values using the less function.
// less: returns whether a < b
func MinByFunc[T any](less func(a, b T) bool, values ...T) T {
	if len(values) == 0 {
		var zero T
		return zero
	}

	min := values[0]
	for i := 1; i < len(values); i++ {
		if less(values[i], min) {
			min = values[i]
		}
	}
	return min
}

// Max returns the maximum element of values
func Max[T cmp.Ordered](values ...T) T {
	return MaxByFunc(func(a, b T) bool { return a < b }, values...)
}

// MaxByFunc returns the maximum element of values using the less function.
// less: returns whether a < b
func MaxByFunc[T any](less func(a, b T) bool, values ...T) T {
	if len(values) == 0 {
		var zero T
		return zero
	}

	max := values[0]
	for i := 1; i < len(values); i++ {
		if less(max, values[i]) {
			max = values[i]
		}
	}
	return max
}

// Unique returns all unique elements in s
func Unique[K comparable](s []K) (u []K) {
	existingElements := map[K]bool{}
	for _, v := range s {
		if existingElements[v] {
			continue
		}
		u = append(u, v)
		existingElements[v] = true
	}
	return
}

// Coalesce, just like the sql function, returns the first non-zero value from values,
// if all values are zero, then the function returns the zero value
func Coalesce[T comparable](values ...T) T {
	var zeroValue T
	for _, v := range values {
		if v != zeroValue {
			return v
		}
	}
	return zeroValue
}

// CoalesceByFunc, just like the sql function, returns the first non-zero value from values,
// it uses isEqual(T, T) to determine whether the values are equals to the zero value,
// if all values are zero, then the function returns the zero value
func CoalesceByFunc[T any](isEqual func(T, T) bool, values ...T) T {
	var zeroValue T
	for _, v := range values {
		if !isEqual(v, zeroValue) {
			return v
		}
	}
	return zeroValue
}

// CompareSlice determines whether 's' and 'target' are equals
// (same elements at the same position and same slice lengths)
func CompareSlice[T comparable](s, target []T) bool {
	if len(s) != len(target) {
		return false
	}
	for i := range s {
		if s[i] != target[i] {
			return false
		}
	}
	return true
}

// IsSubset returns whether a is subset of b (duplicates ignored)
func IsSubset[T comparable](a, b []T) bool {
	if len(a) == 0 {
		return true
	}
	if len(b) == 0 {
		return false
	}
	
	bMap := SliceToBoolMap(b)
	for _, v := range a {
		if !bMap[v] {
			return false
		}
	}
	return true
}