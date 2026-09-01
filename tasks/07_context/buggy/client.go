// ТИКЕТ #1388
// "Сервис после недели работы съедает всю память, перезапускается по OOM.
// Утечка растёт пропорционально числу обработанных запросов."
//
// go vet этот баг НЕ ловит. Профиль pprof показывает рост в runtime, не в Fetch.
package buggy

import (
	"context"
	"io"
	"net/http"
	"time"
)

// Fetch забирает тело по URL с ограничением в 3 секунды.
func Fetch(parent context.Context, url string) ([]byte, error) {
	ctx, _ := context.WithTimeout(parent, 3*time.Second)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(resp.Body)
}
