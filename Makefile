.PHONY: test race lint fmt vet all

# Прогнать все тесты. Красные — это нормально: значит задача ещё не решена.
test:
	go test ./...

# То же с детектором гонок. Обязательно для tasks/05, 06, 07.
race:
	go test -race ./...

# Отформатировать всё по канону.
fmt:
	gofmt -w .

# Статический анализ: ловит то, что компилятор пропускает.
vet:
	go vet ./...

lint: fmt vet

all: lint test
