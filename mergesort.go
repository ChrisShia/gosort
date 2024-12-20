package gosort

//FIRST IMPLEMENTATION

func MergeSortInPlace[T ordered](arr []T) {
	length := len(arr)
	if length <= 1 {
		return
	}
	middle := len(arr) / 2
	leftList := arr[:middle]
	rightList := arr[middle:]
	MergeSortInPlace[T](leftList)
	MergeSortInPlace[T](rightList)
	mergeOrderedContiguousSlicesWithTheSameUnderliningArray(leftList, rightList)
	return
}

func mergeOrderedContiguousSlicesWithTheSameUnderliningArray[T ordered](leftList, rightList []T) {
	for len(leftList) > 0 && len(rightList) > 0 {
		if leftList[0] > rightList[0] {
			if len(rightList) > 1 {
				rightList = rightList[1:]
			} else {
				rightList = rightList[len(rightList):]
			}
			leftList = leftList[:len(leftList)+1]
			rotateRightOnceWithCopy(leftList)
		}
		leftList = leftList[1:]
	}
	return
}

func rotateUsingAuxSlice(s []int, shift int) {
	tmp := make([]int, 0)
	length := len(s)
	noOfElementsToBeRotated := length - shift
	tmp = append(tmp, s[noOfElementsToBeRotated:]...)
	tmp = append(tmp, s[:noOfElementsToBeRotated]...)
	copy(s, tmp)
}

func rotateRightOnceWithCopy[T any](s []T) {
	v := s[len(s)-1]
	copy(s[1:], s[0:len(s)-1])
	s[0] = v
}

//SECOND IMPLEMENTATION

func MergeSortSimple[T ordered](arr []T) []T {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	middle := len(arr) / 2
	sortedLeft1 := MergeSortSimple[T](arr[:middle])
	sortedRight2 := MergeSortSimple[T](arr[middle:])
	return mergeSimple(sortedLeft1, sortedRight2)
}

func mergeSimple[T ordered](left, right []T) []T {
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
