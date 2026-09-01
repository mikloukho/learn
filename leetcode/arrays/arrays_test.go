package arrays

import (
	"reflect"
	"sort"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want bool
	}{
		{"есть повтор", []int{1, 2, 3, 1}, true},
		{"все разные", []int{1, 2, 3, 4}, false},
		{"пустой", []int{}, false},
		{"один элемент", []int{7}, false},
		{"все одинаковые", []int{2, 2, 2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsDuplicate(tt.in); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"базовый", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"в середине", []int{3, 2, 4}, 6, []int{1, 2}},
		{"одинаковые", []int{3, 3}, 6, []int{0, 1}},
		{"отрицательные", []int{-1, -2, -3, -4}, -6, []int{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"обычный", []int{7, 1, 5, 3, 6, 4}, 5},
		{"только падение", []int{7, 6, 4, 3, 1}, 0},
		{"один день", []int{5}, 0},
		{"пустой", []int{}, 0},
		{"рост в конце", []int{2, 4, 1, 7}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProfit(tt.prices); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"базовый", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"один ноль", []int{0}, []int{0}},
		{"нулей нет", []int{1, 2, 3}, []int{1, 2, 3}},
		{"все нули", []int{0, 0, 0}, []int{0, 0, 0}},
		{"ноль в начале", []int{0, 0, 1}, []int{1, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MoveZeroes(tt.in)
			if !reflect.DeepEqual(tt.in, tt.want) {
				t.Errorf("got %v, want %v", tt.in, tt.want)
			}
		})
	}
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name  string
		in    []int
		wantK int
		want  []int
	}{
		{"базовый", []int{1, 1, 2}, 2, []int{1, 2}},
		{"длинный", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
		{"без дублей", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"один", []int{1}, 1, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := RemoveDuplicates(tt.in)
			if k != tt.wantK {
				t.Fatalf("k = %d, want %d", k, tt.wantK)
			}
			if !reflect.DeepEqual(tt.in[:k], tt.want) {
				t.Errorf("got %v, want %v", tt.in[:k], tt.want)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{"базовый", []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{"второй пустой", []int{1}, 1, []int{}, 0, []int{1}},
		{"первый пустой", []int{0}, 0, []int{1}, 1, []int{1}},
		{"второй весь меньше", []int{4, 5, 0, 0}, 2, []int{1, 2}, 2, []int{1, 2, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Merge(tt.nums1, tt.m, tt.nums2, tt.n)
			if !reflect.DeepEqual(tt.nums1, tt.want) {
				t.Errorf("got %v, want %v", tt.nums1, tt.want)
			}
		})
	}
}
