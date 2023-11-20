package int_recursive_merge

func RecursiveMerge(data []int) []int {
	if len(data) < 1 {
		return []int{}
	}

	if len(data) == 1 {
		return data
	}

	if len(data) == 2 {
		if data[0] >= data[1] {
			return []int{data[1], data[0]}
		} else {
			return data
		}
	}

	c := len(data) / 2
	return merge(RecursiveMerge(data[:c]), RecursiveMerge(data[c:]))
}

func merge(data1 []int, data2 []int) []int {
	countSum := len(data1) + len(data2)
	i1 := 0
	i2 := 0
	result := make([]int, 0, countSum)
	for i := 0; i < countSum; i++ {
		if i1 == len(data1) {
			result = append(result, data2[i2:]...)
			break
		}

		if i2 == len(data2) {
			result = append(result, data1[i1:]...)
			break
		}

		if data1[i1] <= data2[i2] {
			result = append(result, data1[i1])
			i1++
		} else {
			result = append(result, data2[i2])
			i2++
		}
	}
	return result
}
