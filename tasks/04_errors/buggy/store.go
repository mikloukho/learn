// ТИКЕТ #1042
// "Функция изредка сообщает об ошибке, хотя данные сохранились корректно.
// Воспроизводится только когда база пустая."
//
// Найди причину, объясни механизм, почини. Тест, доказывающий фикс, — обязателен.
package buggy

import "errors"

type NotFoundError struct{ ID int }

func (e *NotFoundError) Error() string { return "not found" }

type Store struct{ data map[int]string }

func NewStore() *Store { return &Store{data: map[int]string{}} }

// Save сохраняет значение и возвращает ошибку, если запись не удалась.
func (s *Store) Save(id int, v string) error {
	var err *NotFoundError

	if v == "" {
		return errors.New("empty value")
	}
	s.data[id] = v

	return err
}
