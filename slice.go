package goutil

// SliceToMap 泛型函数，将一个包含键值对的切片转换为映射
func SliceToMap[K comparable, V any](items []V, keyFunc func(V) K) map[K]V {
	dictionary := make(map[K]V)

	for _, item := range items {
		dictionary[keyFunc(item)] = item
	}

	return dictionary
}
