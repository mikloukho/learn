// Блок 04. Ошибки, defer, паника.
//
// КОРОТКО: panic — не замена исключению, а "программа в невозможном состоянии".
// Паниковать вместо возврата error — главный маркер новичка из PHP/Java.
//
// ИНСАЙТ (главный в блоке): nil, который не nil.
// Интерфейс внутри — пара указателей (тип, данные). error — интерфейс.
// Положишь в него nil-указатель конкретного типа — пара будет (*MyErr, nil),
// то есть интерфейс НЕ nil, и `if err != nil` сработает при отсутствии ошибки.
//
//	func bad() error { var e *MyErr = nil; return e }  // вернётся НЕ nil!
//
// Правило: не объявляй переменную конкретного типа ошибки заранее, возвращай nil.
//
// СОБЕС:
//   - Почему в Go нет исключений? Аргументируй позицию авторов.
//   - errors.Is vs errors.As vs == — когда что?
//   - Что делает %w и чем отличается от %v?
//   - Может ли error быть != nil, когда внутри nil?
//   - Где работает recover и что будет, если вызвать его не в defer?
//   - Может ли defer изменить возвращаемое значение?
package errorsblock

// ErrNotFound — сентинельная ошибка. Проверяется через errors.Is,
// работает сквозь любое число обёрток.
// TODO: объяви через errors.New.

type Config struct{}

// ValidationError — свой тип ошибки. Достаётся из цепочки через errors.As.
type ValidationError struct {
	Field  string
	Reason string
}

// Error делает тип ошибкой. Один метод — и всё.
func (e *ValidationError) Error() string {
	panic("TODO")
}

// ParseConfig — оберни ошибку чтения файла через fmt.Errorf("read config: %w", err).
// Глагол %w вкладывает ошибку внутрь, %v просто печатает её текст и теряет тип.
func ParseConfig(path string) (*Config, error) {
	panic("TODO")
}

// Validate — верни *ValidationError, обёрнутый через %w. Потом достань его
// в тесте через errors.As и проверь поле Field.
func Validate(email string) error {
	panic("TODO")
}

// Find — верни ошибку, обёрнутую поверх ErrNotFound. В тесте проверь errors.Is.
func Find(id int) error {
	panic("TODO")
}

// SafeDivide — перехвати панику через recover и преврати в error.
// recover работает ТОЛЬКО внутри defer и только на один уровень.
func SafeDivide(a, b int) (res int, err error) {
	panic("TODO")
}

// DeferMutatesResult — сделай так, чтобы defer изменил именованное
// возвращаемое значение уже ПОСЛЕ return. Понять этот трюк = понять,
// как устроен возврат в Go: return присваивает переменную, потом бегут defer.
func DeferMutatesResult() (result string) {
	panic("TODO")
}
