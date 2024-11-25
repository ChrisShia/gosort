package gosort

import (
	"reflect"
	"testing"
)

func TestMergeSortInts(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "",
			input: []int{1, 5, 3, 14, 67, 34, 2, 4},
			want:  []int{1, 2, 3, 4, 5, 14, 34, 67},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := MergeSort(tt.input)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestMergeInts(t *testing.T) {
	tests := []struct {
		name   string
		s1, s2 []int
		want   []int
	}{
		{
			name: "",
			s1:   []int{1, 3, 6},
			s2:   []int{2, 4, 5},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "",
			s1:   []int{1, 3, 6},
			s2:   []int{2, 4, 5, 8, 8},
			want: []int{1, 2, 3, 4, 5, 6, 8, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := merge(tt.s1, tt.s2)
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("Merge() = %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestRotate(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		rotations int
		want      []int
	}{
		{
			name:      "",
			input:     []int{1, 2, 3, 4, 5, 6},
			rotations: 3,
			want:      []int{4, 5, 6, 1, 2, 3},
		},
		{
			name:      "",
			input:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			rotations: 3,
			want:      []int{8, 9, 10, 1, 2, 3, 4, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotateUsingAuxSlice(tt.input, tt.rotations)
			if !reflect.DeepEqual(tt.input, tt.want) {
				t.Errorf("Merge() = %v, want %v", tt.input, tt.want)
			}
		})
	}
}

func TestCopy(t *testing.T) {
	tests := []struct {
		name string
		dst  []int
		src  []int
	}{
		{
			name: "",
			dst:  []int{1, 2, 3, 4, 5, 6},
			src:  []int{8, 9, 10},
		},
		{
			name: "",
			dst:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			src:  []int{1, 2, 3, 4, 5, 11, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			copy(tt.dst[2:], tt.src)
		})
	}
}

func TestRotateOnce(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"", []int{1, 2, 3, 4, 5}, []int{5, 1, 2, 3, 4}},
		{"", []int{}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotateRightOnce(tt.input)
			if !reflect.DeepEqual(tt.input, tt.want) {
				t.Errorf("Merge() = %v, want %v", tt.input, tt.want)
			}
		})
	}
}

func TestMergeWithoutAuxSpace(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{"", []int{1, 3, 4, 9, 2, 5, 7, 11}, []int{1, 2, 3, 4, 5, 7, 9, 11}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			middle := len(tt.arr) / 2
			mergeOrderedContiguousSlicesWithTheSameUnderliningArray(tt.arr[:middle], tt.arr[middle:])
			if !reflect.DeepEqual(tt.arr, tt.want) {
				t.Errorf("Merge() = %v, want %v", tt.arr, tt.want)
			}
		})
	}
}

func TestMergeSortInPlace(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{"", []int{1, 3, 4, 9, 2, 5, 7, 11}, []int{1, 2, 3, 4, 5, 7, 9, 11}},
		{"", []int{1, 10, 15, 3, 100, 0, 7, 11, 90, 56}, []int{0, 1, 3, 7, 10, 11, 15, 56, 90, 100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSortInPlace(tt.arr)
			if !reflect.DeepEqual(tt.arr, tt.want) {
				t.Errorf("Merge() = %v, want %v", tt.arr, tt.want)
			}
		})
	}
}
