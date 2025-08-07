package util

func Keys[K comparable, V any](m map[K]V) (keys []K) {
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
