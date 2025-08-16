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

// Uniques returns all unique elements in s
func Uniques[K comparable](s []K) (u []K) {
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