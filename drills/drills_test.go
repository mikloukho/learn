package drills

import (
	"reflect"
	"testing"
)

func TestModifyElement(t *testing.T) {
	s := []int{1, 2, 3}
	ModifyElement(s)
	if s[0] == 1 {
		t.Error("элемент не изменился: слайс делит массив с вызывающим, изменение должно быть видно")
	}
}

func TestAppendInside(t *testing.T) {
	s := []int{1, 2, 3}
	AppendInside(s)
	if len(s) != 3 {
		t.Errorf("len = %d, want 3: len лежит в КОПИИ заголовка, снаружи он прежний", len(s))
	}
}

func TestSubslice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	sub := Subslice(s)
	if !reflect.DeepEqual(sub, []int{2, 3}) {
		t.Fatalf("got %v, want [2 3]", sub)
	}
	sub[0] = 99
	if s[1] != 99 {
		t.Error("оригинал не изменился: срез обязан делить память с исходным массивом")
	}
}

func TestIndependentCopy(t *testing.T) {
	s := []int{1, 2, 3}
	c := IndependentCopy(s)
	if !reflect.DeepEqual(c, s) {
		t.Fatalf("got %v, want %v", c, s)
	}
	c[0] = 99
	if s[0] == 99 {
		t.Error("оригинал изменился: нужна настоящая копия через make + copy")
	}
}

func TestLenCap(t *testing.T) {
	l, c := LenCap()
	if l != 2 || c != 10 {
		t.Errorf("got len=%d cap=%d, want len=2 cap=10", l, c)
	}
}

func TestRemoveLastNoLeak(t *testing.T) {
	a, b, c := 1, 2, 3
	s := []*int{&a, &b, &c}
	got := RemoveLastNoLeak(s)
	if len(got) != 2 {
		t.Fatalf("len = %d, want 2", len(got))
	}
	// Заглядываем за длину — там обязан лежать nil, иначе объект держится в памяти.
	if s[:3][2] != nil {
		t.Error("хвост не обнулён: указатель остался в массиве и GC не заберёт объект")
	}
}

func TestModifyArray(t *testing.T) {
	a := [3]int{1, 2, 3}
	ModifyArray(a)
	if a[0] != 1 {
		t.Error("массив изменился снаружи: он копируется целиком, так быть не должно")
	}
}

func TestArraysEqual(t *testing.T) {
	if !ArraysEqual([3]int{1, 2, 3}, [3]int{1, 2, 3}) {
		t.Error("одинаковые массивы должны быть равны")
	}
	if ArraysEqual([3]int{1, 2, 3}, [3]int{1, 2, 4}) {
		t.Error("разные массивы не должны быть равны")
	}
}

func TestMapLookup(t *testing.T) {
	m := map[string]int{"a": 0}
	if v, ok := MapLookup(m, "a"); v != 0 || !ok {
		t.Errorf("got (%d, %v), want (0, true): ключ есть, значение нулевое", v, ok)
	}
	if v, ok := MapLookup(m, "нет"); v != 0 || ok {
		t.Errorf("got (%d, %v), want (0, false)", v, ok)
	}
}

func TestAddToMap(t *testing.T) {
	m := map[string]int{}
	AddToMap(m, "k", 7)
	if m["k"] != 7 {
		t.Error("ключ не появился: мапа ссылочная, изменение видно снаружи")
	}
}

func TestReadNilMap(t *testing.T) {
	var m map[string]int // nil
	if got := ReadNilMap(m, "k"); got != 0 {
		t.Errorf("got %d, want 0", got)
	}
}

func TestSortedKeys(t *testing.T) {
	m := map[string]int{"b": 1, "a": 2, "c": 3}
	if got, want := SortedKeys(m), []string{"a", "b", "c"}; !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestByteAndRuneLen(t *testing.T) {
	if b, r := ByteAndRuneLen("привет"); b != 12 || r != 6 {
		t.Errorf("got bytes=%d runes=%d, want 12 и 6", b, r)
	}
	if b, r := ByteAndRuneLen("hi"); b != 2 || r != 2 {
		t.Errorf("got bytes=%d runes=%d, want 2 и 2", b, r)
	}
}

func TestReplaceFirstByte(t *testing.T) {
	if got := ReplaceFirstByte("cat", 'b'); got != "bat" {
		t.Errorf("got %q, want \"bat\"", got)
	}
}

func TestFirstByte(t *testing.T) {
	if got := FirstByte("abc"); got != 'a' {
		t.Errorf("got %d, want %d", got, 'a')
	}
	if got := FirstByte("привет"); got != 208 {
		t.Errorf("got %d, want 208 — первый байт двухбайтовой буквы", got)
	}
}

func TestFirstRune(t *testing.T) {
	if got := FirstRune("привет"); got != 'п' {
		t.Errorf("got %q, want 'п'", got)
	}
}
