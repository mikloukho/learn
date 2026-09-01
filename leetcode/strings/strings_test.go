package strings

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		name string
		in   []byte
		want []byte
	}{
		{"чётная длина", []byte("abcd"), []byte("dcba")},
		{"нечётная", []byte("hello"), []byte("olleh")},
		{"один символ", []byte("x"), []byte("x")},
		{"пустая", []byte(""), []byte("")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseString(tt.in)
			if !reflect.DeepEqual(tt.in, tt.want) {
				t.Errorf("got %q, want %q", tt.in, tt.want)
			}
		})
	}
}

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s, t string
		want bool
	}{
		{"анаграмма", "anagram", "nagaram", true},
		{"не анаграмма", "rat", "car", false},
		{"разная длина", "a", "ab", false},
		{"обе пустые", "", "", true},
		{"те же буквы, разное число", "aab", "abb", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAnagram(tt.s, tt.t); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{"с пунктуацией", "A man, a plan, a canal: Panama", true},
		{"не палиндром", "race a car", false},
		{"пробел", " ", true},
		{"пустая", "", true},
		{"цифры", "0P", false},
		{"только цифры", "12321", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.in); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		want string
	}{
		{"есть префикс", []string{"flower", "flow", "flight"}, "fl"},
		{"нет общего", []string{"dog", "racecar", "car"}, ""},
		{"одна строка", []string{"alone"}, "alone"},
		{"пустая среди прочих", []string{"abc", ""}, ""},
		{"полное совпадение", []string{"same", "same"}, "same"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonPrefix(tt.in); got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestIsValidParentheses(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{"простые", "()", true},
		{"все виды", "()[]{}", true},
		{"неверный порядок", "([)]", false},
		{"вложенные", "{[()]}", true},
		{"незакрытая", "(", false},
		{"лишняя закрывающая", "){", false},
		{"пустая", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidParentheses(tt.in); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseUnicode(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"кириллица", "привет", "тевирп"},
		{"латиница", "hello", "olleh"},
		{"пустая", "", ""},
		{"эмодзи", "aб🙂", "🙂бa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseUnicode(tt.in); got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
