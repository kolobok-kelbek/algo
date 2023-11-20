package compare_recursive_merge

type CompareFunc[T comparable] func(value1 T, value2 T) bool

func RecursiveMerge[T comparable](compare CompareFunc[T], data []T) []T {
	if len(data) < 1 {
		return []T{}
	}

	if len(data) == 1 {
		return data
	}

	if len(data) == 2 {
		if compare(data[0], data[1]) {
			return []T{data[1], data[0]}
		} else {
			return data
		}
	}

	c := len(data) / 2
	return merge(compare, RecursiveMerge(compare, data[:c]), RecursiveMerge(compare, data[c:]))
}

func merge[T comparable](compare CompareFunc[T], data1 []T, data2 []T) []T {
	countSum := len(data1) + len(data2)
	i1 := 0
	i2 := 0
	result := make([]T, 0, countSum)
	for i := 0; i < countSum; i++ {
		if i1 == len(data1) {
			result = append(result, data2[i2:]...)
			break
		}

		if i2 == len(data2) {
			result = append(result, data1[i1:]...)
			break
		}

		if compare(data1[i1], data2[i2]) {
			result = append(result, data2[i2])
			i2++
		} else {
			result = append(result, data1[i1])
			i1++
		}
	}
	return result
}
