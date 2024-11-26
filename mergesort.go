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

func MergeSortInPlaceParallel[T ordered](arr []T, done chan<- bool) {
	length := len(arr)
	if length <= 1 {
		return
	}
	middle := len(arr) / 2
	leftList := arr[:middle]
	rightList := arr[middle:]
	if length >= 10 && done != nil {
		for d := range mergeSortInPlaceParallel[T](leftList, rightList) {
			done <- d
		}
		close(done)
		return
	}
	MergeSortInPlace[T](leftList)
	MergeSortInPlace[T](rightList)
	mergeOrderedContiguousSlicesWithTheSameUnderliningArray(leftList, rightList)
	done <- true
	close(done)
	return
}

func mergeSortInPlaceParallel[T ordered](leftList, rightList []T) <-chan bool {
	done := make(chan bool)
	go func() {
		doneLeft := make(chan bool)
		doneRight := make(chan bool)
		go MergeSortInPlaceParallel[T](leftList, doneLeft)
		go MergeSortInPlaceParallel[T](rightList, doneRight)
		//doneWithLeft := <-doneLeft
		//doneWithRight := <-doneRight
		if <-doneLeft && <-doneRight {
			mergeOrderedContiguousSlicesWithTheSameUnderliningArray[T](leftList, rightList)
			done <- true
			close(done)
		}
	}()
	return done
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
			rotateRightOnce(leftList)
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

func rotateRightOnce[T any](s []T) {
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
