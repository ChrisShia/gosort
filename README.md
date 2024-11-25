The provided code consists of two implementations of the Merge Sort algorithm in Go. Merge Sort is a classic divide-and-conquer algorithm used to sort elements. The main difference between the two implementations lies in how they handle the sorting process in place versus returning a new sorted slice.

### Package Declaration
```go
package gosort
```
This declares that the code is part of the `gosort` package.

### FIRST IMPLEMENTATION (PREFERRED)

#### Function: `MergeSortInPlace`
```go
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
```
- **Purpose**: This function sorts a slice of elements in place.
- **Parameters**: 
  - `arr`: A slice of any ordered type (e.g., `int`, `float64`, `string`).
- **Steps**:
  1. If the length of the slice is 1 or less, it returns immediately as the slice is already sorted.
  2. It splits the slice into two halves: `leftList` and `rightList`.
  3. It recursively sorts `leftList` and `rightList` using the same function.
  4. It merges the two sorted halves back into the original slice using the helper function `mergeOrderedContiguousSlicesWithTheSameUnderliningArray`.

#### Function: `mergeOrderedContiguousSlicesWithTheSameUnderliningArray`
```go
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
```
- **Purpose**: This function merges two already sorted slices back into the original slice using a single underlying array.
- **Parameters**:
  - `leftList`: The first sorted slice.
  - `rightList`: The second sorted slice.
- **Steps**:
  1. It iterates through both slices while both have elements.
  2. If the first element of `leftList` is greater than the first element of `rightList`, it removes the first element of `rightList` and appends it to `leftList`. It then rotates `leftList` to the right by one position.
  3. It continues this process until one of the slices is empty.
- **Note**: This implementation uses `rotateRightOnce` to shift elements within the slice, which might be inefficient for large slices.

#### Function: `rotateRightOnce`
```go
func rotateRightOnce[T any](s []T) {
    v := s[len(s)-1]
    copy(s[1:], s[0:len(s)-1])
    s[0] = v
}
```
- **Purpose**: This function rotates a slice to the right by one position.
- **Parameters**:
  - `s`: A slice of any type.
- **Steps**:
  1. It stores the last element of the slice in a temporary variable `v`.
  2. It shifts all elements to the right by one position.
  3. It assigns the stored last element to the first position of the slice.

### SECOND IMPLEMENTATION

#### Function: `MergeSort`
```go
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
```
- **Purpose**: This function sorts a slice of elements and returns a new sorted slice.
- **Parameters**:
  - `arr`: A slice of any ordered type.
- **Steps**:
  1. If the length of the slice is 1 or less, it returns the slice itself as it is already sorted.
  2. It splits the slice into two halves: `leftList` and `rightList`.
  3. It recursively sorts `leftList` and `rightList` and assigns the results to `sortedLeft1` and `sortedRight2`.
  4. It merges the two sorted slices using the helper function `merge`.

#### Function: `merge`
```go
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
```
- **Purpose**: This function merges two sorted slices into a single sorted slice.
- **Parameters**:
  - `left`: The first sorted slice.
  - `right`: The second sorted slice.
- **Steps**:
  1. It creates a new slice `merged` of length equal to the sum of the lengths of `left` and `right`.
  2. It uses three counters: `i` for `left`, `j` for `right`, and `m` for `merged`.
  3. It iterates through both slices, comparing elements and appending the smaller element to `merged`.
  4. After one of the slices is exhausted, it appends the remaining elements of the other slice to `merged`.
  5. It returns the merged slice.

### Summary
- **FIRST IMPLEMENTATION**: Sorts the slice in place using a single underlying array for merging. It uses a helper function to rotate elements within the slice, which might be inefficient.
- **SECOND IMPLEMENTATION**: Returns a new sorted slice by merging two sorted halves.

Both implementations achieve the same result but differ in terms of efficiency and in-place sorting capability. The second implementation is generally preferred for its simplicity and ease of understanding.

