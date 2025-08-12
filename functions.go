package generics

import (
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