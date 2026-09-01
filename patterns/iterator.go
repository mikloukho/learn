package patterns

// Iterator.
// В Go до 1.23 итератор — это интерфейс HasNext()/Next()
// или канал. Начиная с 1.23 — функции-итераторы (range over func).
//
// TODO: интерфейс Iterator, коллекция, метод CreateIterator().
