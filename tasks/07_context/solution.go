// Блок 07. Context.
//
// КОРОТКО: ctx — всегда ПЕРВЫЙ аргумент, называется ctx, НИКОГДА не хранится
// в структуре, не передаётся nil (для заглушки есть context.TODO()).
//
// ИНСАЙТ: context отменяет не выполнение, а ОЖИДАНИЕ. Отмена не убивает
// горутину, она закрывает канал Done. Не проверяешь его — работа продолжится
// как ни в чём не бывало. Убить горутину снаружи в Go нельзя вообще.
//
// ИНСАЙТ: cancel надо вызывать ВСЕГДА, даже при успешном завершении.
// Иначе родительский контекст держит ссылку на дочерний — это утечка памяти.
//
// СОБЕС:
//   - Какие есть виды контекста? Почему ctx нельзя хранить в структуре?
//   - Что будет, если не вызвать cancel?
//   - Отменяет ли ctx выполнение горутины?
//   - Когда уместен WithValue, а когда нет?
//   - Как сделать graceful shutdown HTTP-сервера?
package contextblock

import "context"

// LongWork — долгая работа, прерываемая через <-ctx.Done() в select.
// Верни ctx.Err(), если отменили.
func LongWork(ctx context.Context) error {
	panic("TODO")
}

// FetchWithTimeout — запрос, обязанный уложиться в 2 секунды.
// context.WithTimeout + defer cancel().
func FetchWithTimeout(ctx context.Context, url string) ([]byte, error) {
	panic("TODO")
}

// Chain — цепочка из трёх вызовов, ctx пробрасывается насквозь.
// Отмена наверху должна гасить всю цепочку.
func Chain(ctx context.Context) error {
	panic("TODO")
}

// WithRequestID / RequestID — прокинуть request-id через WithValue.
// Объясни, почему класть туда параметры бизнес-логики — плохая практика
// (подсказка: теряется типобезопасность и неявные зависимости).
func WithRequestID(ctx context.Context, id string) context.Context {
	panic("TODO")
}

func RequestID(ctx context.Context) (string, bool) {
	panic("TODO")
}

// RunServer — graceful shutdown: ловим SIGTERM через signal.NotifyContext,
// даём серверу 5 секунд дожить через srv.Shutdown(ctx).
func RunServer(addr string) error {
	panic("TODO")
}
