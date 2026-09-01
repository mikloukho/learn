package datastructures

import "testing"

// values — хелпер для тестов: собирает список в слайс.
func (l *List) values() []int {
	var out []int
	for curr := l.head; curr != nil; curr = curr.next {
		out = append(out, curr.data)
	}
	return out
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestAdd(t *testing.T) {
	l := &List{}
	for _, v := range []int{1, 2, 3} {
		l.add(v)
	}
	if got, want := l.values(), []int{1, 2, 3}; !equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name   string
		add    []int
		remove int
		want   []int
	}{
		{"из середины", []int{1, 2, 3, 4}, 2, []int{1, 3, 4}},
		{"голову", []int{1, 2, 3}, 1, []int{2, 3}},
		{"хвост", []int{1, 2, 3}, 3, []int{1, 2}},
		{"единственный", []int{5}, 5, nil},
		{"которого нет", []int{1, 2}, 9, []int{1, 2}},
		{"из пустого", nil, 1, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{}
			for _, v := range tt.add {
				l.add(v)
			}
			l.remove(tt.remove)
			if got := l.values(); !equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
