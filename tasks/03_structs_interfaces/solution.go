// Блок 03. Структуры и интерфейсы.
//
// ИНСАЙТ: интерфейс внутри — ПАРА указателей (таблица типа, данные).
// Отсюда: интерфейс с nil-указателем внутри сам НЕ nil. Подробнее в блоке 04.
//
// ИНСАЙТ: struct{} занимает 0 байт — поэтому это значение для множеств.
//
// ИНСАЙТ: порядок полей влияет на размер структуры из-за выравнивания.
// bool,int64,bool = 24 байта, а bool,bool,int64 = 16.
//
// ПРАВИЛА: интерфейс объявляется у ПОТРЕБИТЕЛЯ и он маленький (1-3 метода).
// Принимай интерфейсы, возвращай структуры. Экспортируется всё с Большой буквы.
//
// СОБЕС:
//   - Value receiver vs pointer receiver — когда что?
//   - Почему *T реализует интерфейс, а T нет, если метод на pointer-receiver?
//   - Type assertion vs type switch, что вернёт v, ok := i.(string)?
//   - Встраивание — это наследование? (нет, объясни разницу)
//   - Пустой интерфейс any — чем опасен?
package structs

type User struct {
	ID    int
	Name  string
	Email string
}

// NewUser — конструктор это просто функция, а не метод. Валидируй email.
func NewUser(id int, name, email string) (*User, error) {
	panic("TODO")
}

// String реализует fmt.Stringer НЕЯВНО — никаких implements.
// Проверь: fmt.Println(u) сам его вызовет.
func (u User) String() string {
	panic("TODO")
}

// Rename — pointer receiver. Объясни, почему на value-receiver он бы не работал.
func (u *User) Rename(name string) {
	panic("TODO")
}

// Storage — маленький интерфейс. Две реализации ниже.
type Storage interface {
	Save(u *User) error
	Get(id int) (*User, error)
}

// MemoryStorage — хранит в мапе.
type MemoryStorage struct{}

// NullStorage — ничего не делает, всегда успех. Нужен для тестов.
type NullStorage struct{}

// TODO: реализуй Storage для обоих типов.

// Register работает с интерфейсом, а не с конкретным хранилищем.
func Register(s Storage, name, email string) error {
	panic("TODO")
}

// Admin — ВСТРАИВАНИЕ, не наследование. Покажи, что admin.Name работает,
// и объясни, почему Admin при этом не является User с точки зрения типов.
type Admin struct {
	User
	Level int
}
