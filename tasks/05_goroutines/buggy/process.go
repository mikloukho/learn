// ТИКЕТ #1103
// "Обработчик пачки задач иногда возвращает не все результаты, иногда виснет.
// Чем больше задач, тем чаще."
//
// Здесь БОЛЬШЕ ОДНОГО бага. Найди все, объясни каждый, почини.
// Не запускай go vet, пока не поищешь глазами — он назовёт один из них прямо.
package buggy

import "sync"

// Process обрабатывает задачи параллельно и собирает результаты.
func Process(tasks []int) []int {
	results := make(chan int)
	var wg sync.WaitGroup

	for _, t := range tasks {
		wg.Add(1)
		go func() {
			defer wg.Done()
			results <- t * 2
		}()
	}

	wg.Wait()
	close(results)

	var out []int
	for r := range results {
		out = append(out, r)
	}

	return out
}
