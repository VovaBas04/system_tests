package api_gateway

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type Proxy struct {
	config ConfigUrl
}

func NewProxy(config ConfigUrl) *Proxy {
	return &Proxy{config: config}
}

const (
	CodeChecker = iota
	Tasks
)

var NotFoundError = errors.New("not found path")

var routes = map[string]int{
	"/check":  CodeChecker,
	"/tasks":  Tasks,
	"/models": Tasks,
}

func (proxy *Proxy) proxyHandler(c *gin.Context) {
	targetURL, err := proxy.getUrlByPath(c.Request.URL.Path)
	fmt.Println(targetURL)

	if err != nil {
		c.String(http.StatusNotFound, "Ошибка не найдена")
		return
	}

	// Создаем новый запрос к целевому серверу

	retry := NewRetry()
	client := &http.Client{}
	ByteBody, _ := io.ReadAll(c.Request.Body)
	resp := retry.Exec(func() (*http.Response, bool) {
		req, err := http.NewRequest(c.Request.Method, targetURL, io.NopCloser(bytes.NewBuffer(ByteBody)))
		if err != nil {
			fmt.Println(err, "Req")
			return nil, false
		}

		// Копируем заголовки оригинального запроса
		for key, value := range c.Request.Header {
			req.Header[key] = value
		}

		response, err := client.Do(req)

		if err != nil {
			fmt.Println(err, "Resp")
			return nil, false
		}

		if response.StatusCode == http.StatusInternalServerError {
			return response, true
		}

		return response, false
	})
	if resp == nil {
		c.String(http.StatusServiceUnavailable, "Ошибка доступа")
		return
	}
	defer resp.Body.Close()

	// Читаем ответ от целевого сервера
	body, _ := io.ReadAll(resp.Body)

	// Устанавливаем статус код и заголовки из ответа
	c.Writer.WriteHeader(resp.StatusCode)
	for key, value := range resp.Header {
		c.Writer.Header()[key] = value
	}

	// Возвращаем ответ обратно клиенту
	c.Writer.Write(body)
}

func (proxy *Proxy) getUrlByPath(path string) (string, error) {
	value, ok := routes[path]
	if !ok {
		return "", NotFoundError
	}

	if value == CodeChecker {
		return proxy.config.GetUrlCodeChecker() + path, nil
	}

	return proxy.config.GetUrlTasks() + path, nil
}
