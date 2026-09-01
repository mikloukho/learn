package lists

import (
	"reflect"
	"testing"
)

// build собирает список из слайса, fromList разбирает обратно.
func build(vals []int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, v := range vals {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}
	return dummy.Next
}

func fromList(head *ListNode) []int {
	var out []int
	for ; head != nil; head = head.Next {
		out = append(out, head.Val)
	}
	return out
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name string
		a, b []int
		want []int
	}{
		{"базовый", []int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{"оба пустые", nil, nil, nil},
		{"первый пустой", nil, []int{0}, []int{0}},
		{"второй пустой", []int{1, 2}, nil, []int{1, 2}},
		{"не пересекаются", []int{1, 2}, []int{8, 9}, []int{1, 2, 8, 9}},
		{"все равны", []int{2, 2}, []int{2}, []int{2, 2, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fromList(MergeTwoLists(build(tt.a), build(tt.b)))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"обычный", []int{1, 2, 3}, []int{3, 2, 1}},
		{"один", []int{1}, []int{1}},
		{"пустой", nil, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fromList(ReverseList(build(tt.in)))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasCycle(t *testing.T) {
	t.Run("без цикла", func(t *testing.T) {
		if HasCycle(build([]int{1, 2, 3})) {
			t.Error("got true, want false")
		}
	})
	t.Run("пустой", func(t *testing.T) {
		if HasCycle(nil) {
			t.Error("got true, want false")
		}
	})
	t.Run("с циклом", func(t *testing.T) {
		head := build([]int{1, 2, 3})
		tail := head
		for tail.Next != nil {
			tail = tail.Next
		}
		tail.Next = head.Next // замыкаем на вторую ноду
		if !HasCycle(head) {
			t.Error("got false, want true")
		}
	})
}
