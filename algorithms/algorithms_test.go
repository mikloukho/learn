package algorithms

import (
	"reflect"
	"testing"
)

func TestEuclidean(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"классика", 270, 192, 6},
		{"взаимно простые", 13, 7, 1},
		{"кратные", 100, 25, 25},
		{"ноль вторым", 5, 0, 5},
		{"равные", 8, 8, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Euclidean(tt.a, tt.b); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name    string
		items   []int
		target  int
		wantIdx int
		wantOK  bool
	}{
		{"в середине", []int{1, 3, 5, 7, 9}, 5, 2, true},
		{"первый", []int{1, 3, 5}, 1, 0, true},
		{"последний", []int{1, 3, 5}, 5, 2, true},
		{"нет такого", []int{1, 3, 5}, 4, 0, false},
		{"пустой", []int{}, 1, 0, false},
		{"один элемент, есть", []int{7}, 7, 0, true},
		{"один элемент, нет", []int{7}, 1, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx, ok := BinarySearch(tt.items, tt.target)
			if ok != tt.wantOK {
				t.Fatalf("ok = %v, want %v", ok, tt.wantOK)
			}
			if ok && idx != tt.wantIdx {
				t.Errorf("idx = %d, want %d", idx, tt.wantIdx)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"обычный", []int{5, 2, 9, 1}, []int{1, 2, 5, 9}},
		{"уже отсортирован", []int{1, 2, 3}, []int{1, 2, 3}},
		{"обратный", []int{3, 2, 1}, []int{1, 2, 3}},
		{"с дублями", []int{2, 1, 2}, []int{1, 2, 2}},
		{"пустой", []int{}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectionSort(tt.in)
			if !reflect.DeepEqual(tt.in, tt.want) {
				t.Errorf("got %v, want %v", tt.in, tt.want)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"обычный", []int{5, 2, 9, 1, 7}, []int{1, 2, 5, 7, 9}},
		{"с дублями", []int{3, 1, 3, 1}, []int{1, 1, 3, 3}},
		{"один", []int{1}, []int{1}},
		{"отрицательные", []int{0, -5, 3}, []int{-5, 0, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBreadthFirstSearch(t *testing.T) {
	// Граф знакомств: ищем продавца манго (имя заканчивается на "m").
	g := Graph{
		"you":    {"alice", "bob", "claire"},
		"alice":  {"peggy"},
		"bob":    {"anuj", "peggy"},
		"claire": {"thom", "jonny"},
		"anuj":   {},
		"peggy":  {},
		"thom":   {},
		"jonny":  {},
	}
	isSeller := func(name string) bool { return len(name) > 0 && name[len(name)-1] == 'm' }

	t.Run("находит", func(t *testing.T) {
		got, ok := BreadthFirstSearch(g, "you", isSeller)
		if !ok || got != "thom" {
			t.Errorf("got %q, %v; want \"thom\", true", got, ok)
		}
	})
	t.Run("не находит", func(t *testing.T) {
		if _, ok := BreadthFirstSearch(g, "you", func(string) bool { return false }); ok {
			t.Error("got true, want false")
		}
	})
	t.Run("цикл в графе не вешает", func(t *testing.T) {
		cyclic := Graph{"a": {"b"}, "b": {"a"}}
		if _, ok := BreadthFirstSearch(cyclic, "a", func(s string) bool { return s == "z" }); ok {
			t.Error("got true, want false")
		}
	})
}
