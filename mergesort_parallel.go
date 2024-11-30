package gosort

//TODO: Need to apply a limit in the spawning of routines, to limit what seems to be significant overhead, needs investigation...

func MergeSortInPlaceParallel[T ordered](arr []T, done chan<- bool) {
	length := len(arr)
	if length <= 1 {
		done <- true
	} else {
		middle := len(arr) / 2
		leftList := arr[:middle]
		rightList := arr[middle:]
		done <- <-mergeSortInPlaceParallel(leftList, rightList)
	}
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
		if <-doneLeft && <-doneRight {
			mergeOrderedContiguousSlicesWithTheSameUnderliningArray[T](leftList, rightList)
			done <- true
			close(done)
		}
		return
	}()
	return done
}
