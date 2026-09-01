// Блок 06. Гонки и синхронизация.
//
// КОРОТКО: гонку нельзя найти чтением кода, она не воспроизводится стабильно.
// ВЕСЬ блок гоняй через `go test -race`, иначе задания бессмысленны.
//
// ИНСАЙТ: мьютекс нельзя копировать. sync.Mutex — структура с состоянием.
// Скопируешь структуру, где он лежит, — получишь два независимых мьютекса,
// защищающих одно. Отсюда правило pointer-receiver для всего с мьютексом.
//
// ИНСАЙТ: конкурентная запись в map даёт не порчу данных, а fatal error:
// concurrent map writes — и его НЕЛЬЗЯ перехватить recover. Рантайм намеренно
// убивает программу: молча испорченная хеш-таблица хуже падения.
//
// ИНСАЙТ: -race не анализирует код, он инструментирует обращения к памяти
// в рантайме. Не выполнилась ветка — гонка не найдена.
//
// СОБЕС:
//   - race condition vs data race — в чём разница?
//   - Mutex vs RWMutex vs atomic — что быстрее и когда что?
//   - Почему нельзя копировать мьютекс?
//   - Как реализован sync.Once? Зачем sync.Pool?
//   - Когда sync.Map вместо map+mutex? (редко: read-heavy, разные ключи)
//   - Что такое happens-before в Go memory model?
package race

// UnsafeCounter — СПЕЦИАЛЬНО без синхронизации.
// Запусти 1000 горутин на Inc, посмотри результат, потом прогони с -race
// и прочитай отчёт целиком. Это упражнение, а не ошибка.
type UnsafeCounter struct{ n int }

func (c *UnsafeCounter) Inc()     { panic("TODO") }
func (c *UnsafeCounter) Get() int { panic("TODO") }

// Counter — то же самое с sync.Mutex. Мьютекс кладётся В структуру, первым полем.
// Методы обязаны быть на pointer-receiver — объясни почему.
type Counter struct{}

// Cache — много читателей, редкие записи. sync.RWMutex.
// RLock допускает много одновременных читателей, Lock — эксклюзивен.
type Cache struct{}

// Config + GetConfig — ленивая инициализация через sync.Once.
// Это и есть Singleton из patterns/singleton.go.
type Config struct{}

func GetConfig() *Config { panic("TODO") }

// AtomicCounter — то же, что Counter, но на sync/atomic.
// Напиши бенчмарк и сравни с мьютексом. Разница будет заметная.
type AtomicCounter struct{}
