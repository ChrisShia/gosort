package gosort

func MergeSort[T ordered](arr []T) []T {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	middle := len(arr) / 2
	sortedLeft1 := MergeSort[T](arr[:middle])
	sortedRight2 := MergeSort[T](arr[middle:])
	return merge(sortedLeft1, sortedRight2)
}

func merge[T ordered](left, right []T) []T {
	merged := make([]T, len(left)+len(right))
	i, j, m := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged[m] = left[i]
			i++
		} else {
			merged[m] = right[j]
			j++
		}
		m++
	}
	if i < len(left) {
		for i < len(left) {
			merged[m] = left[i]
			m++
			i++
		}
	} else if j < len(right) {
		for j < len(right) {
			merged[m] = right[j]
			m++
			j++
		}
	}
	return merged
}

func MergeSortInPlace(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}
	middle := len(arr) / 2
	leftList := arr[:middle]
	rightList := arr[middle:]
	MergeSortInPlace(leftList)
	MergeSortInPlace(rightList)
	mergeOrderedContiguousSlicesWithTheSameUnderliningArray(leftList, rightList)
	return
}

func mergeOrderedContiguousSlicesWithTheSameUnderliningArray(leftList, rightList []int) {
	for len(leftList) > 0 && len(rightList) > 0 {
		if leftList[0] > rightList[0] {
			if len(rightList) > 1 {
				rightList = rightList[1:]
			} else {
				rightList = rightList[len(rightList):]
			}
			leftList = leftList[:len(leftList)+1]
			rotateRightOnce(leftList)
		}
		leftList = leftList[1:]
	}
	return
}

type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

func rotateUsingAuxSlice(s []int, shift int) {
	tmp := make([]int, 0)
	length := len(s)
	noOfElementsToBeRotated := length - shift
	tmp = append(tmp, s[noOfElementsToBeRotated:]...)
	tmp = append(tmp, s[:noOfElementsToBeRotated]...)
	copy(s, tmp)
}

func rotateRightOnce(s []int) {
	if s == nil || len(s) == 0 {
		return
	}
	v := s[len(s)-1]
	copy(s[1:], s[0:len(s)-1])
	s[0] = v
}
