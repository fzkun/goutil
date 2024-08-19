package goutil

// SliceToMap 泛型函数，将一个包含键值对的切片转换为映射
func SliceToMap[K comparable, V any](list []V, keyFunc func(V) K) map[K]V {
	dictionary := make(map[K]V)

	for _, item := range list {
		dictionary[keyFunc(item)] = item
	}

	return dictionary
}

func SliceForEach[S ~[]E, E any](list S, f func(E)) {
	for i := range list {
		f(list[i])
	}
}

// SliceToField 数组对象返回一个字段的数组
func SliceToField[T any, F any](arr []T, field func(T) F) []F {
	var result []F
	for _, item := range arr {
		result = append(result, field(item))
	}
	return result
}

// SliceRemoveDuplicates 定义一个泛型函数，用于去重
func SliceRemoveDuplicates[T comparable](arr []T) []T {
	seen := make(map[T]bool)
	result := []T{}

	for _, value := range arr {
		if !seen[value] {
			seen[value] = true
			result = append(result, value)
		}
	}

	return result
}
