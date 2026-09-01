// ТИКЕТ #1247
// "Счётчик просмотров на проде отстаёт от реальности процентов на пять.
// Локально всё сходится, тесты зелёные."
//
// Мьютекс на месте, Lock/Unlock есть. Почему всё равно гонка?
package buggy

import "sync"

type Views struct {
	mu    sync.Mutex
	count map[string]int
}

func NewViews() *Views { return &Views{count: map[string]int{}} }

func (v *Views) Inc(page string) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.count[page]++
}

// Get возвращает текущее количество просмотров.
func (v Views) Get(page string) int {
	return v.count[page]
}
