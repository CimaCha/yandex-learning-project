package main

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// клиент с контролируемой частотой запросов
type Client struct {
	client      *http.Client
	rateLimiter *rate.Limiter
}

// метод Do для клиента
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	// ограничение частоты
	ctx := context.Background()
	// здесь используем блокирующий вызов Wait, а не игнорирующий Allow
	err := c.rateLimiter.Wait(ctx)
	if err != nil {
		return nil, err
	}
	// используем метод Do встроенного http.Client
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// конструктор для клиента
func NewClient(rl *rate.Limiter) *Client {
	c := &Client{
		client:      http.DefaultClient,
		rateLimiter: rl,
	}
	return c
}

func main() {
	// параметры Limiter: не более 50 запросов за 10 секунд
	rl := rate.NewLimiter(rate.Every(10*time.Second), 50)
	// конструируем клиент
	client := NewClient(rl)
	// дальше в коде ничего менять не понадобится
	URL := "https://iss.moex.com/iss/statistics/engines/futures/markets/indicativerates/securities.xml"
	req, _ := http.NewRequest("GET", URL, nil)
	for {
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// логика обработки результата запроса
		defer resp.Body.Close()
		scanner := bufio.NewScanner(resp.Body)
		for i := 0; scanner.Scan() && i < 20; i++ {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}
}
