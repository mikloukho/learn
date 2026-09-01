// Блок 08. Дженерики.
//
// КОРОТКО: дженерик нужен, когда пишешь ОДИН алгоритм для разных типов.
// В остальных случаях — интерфейс или конкретный тип. Половина stdlib без них.
//
// ИНСАЙТ: Go не генерирует копию кода на каждый тип, как C++. Компилятор
// группирует типы по форме представления в памяти (GC shape stenciling)
// и переиспользует код, передавая словарь типов. Побочный эффект: дженерик-код
// часто МЕДЛЕННЕЕ конкретного, потому что теряет инлайнинг. Сначала бенчмарк.
//
// СОБЕС:
//   - Зачем дженерики, если есть interface{}?
//   - Что означает тильда ~int в constraint?
//   - Что такое comparable и почему не всякий тип comparable?
//   - Дженерики быстрее или медленнее конкретного кода?
package generics

// Number — свой constraint. Объясни, что делает тильда:
// ~int означает "int и любой тип с базовым типом int" (type MyInt int).
type Number interface {
	~int | ~int64 | ~float64
}

func Map[T, U any](in []T, f func(T) U) []U {
	panic("TODO")
}

func Filter[T any](in []T, f func(T) bool) []T {
	panic("TODO")
}

func Reduce[T, U any](in []T, init U, f func(U, T) U) U {
	panic("TODO")
}

// Sum — на своём constraint Number.
func Sum[T Number](items []T) T {
	panic("TODO")
}

// Max — на готовом cmp.Ordered из stdlib.
func Max[T Number](items []T) T {
	panic("TODO")
}

// Unique — тот же, что в блоке 02, но на comparable.
// comparable = тип, к которому применим ==. Слайсы и мапы не comparable.
func Unique[T comparable](in []T) []T {
	panic("TODO")
}

// Stack — дженерик-структура.
type Stack[T any] struct{}

func (s *Stack[T]) Push(v T)       { panic("TODO") }
func (s *Stack[T]) Pop() (T, bool) { panic("TODO") }
func (s *Stack[T]) Len() int       { panic("TODO") }
