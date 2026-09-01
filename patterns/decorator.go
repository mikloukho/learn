package patterns

// Decorator.
// Тип реализует тот же интерфейс и хранит внутри следующий слой.
// Самый "гошный" паттерн: так устроены http.Handler middleware.
//
// TODO: интерфейс Notifier, базовая реализация, декораторы SMS/Slack.
