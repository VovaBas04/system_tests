package api_gateway

import (
	"fmt"
	"log"
	"net/http"
)

const defaultMaxRetries = 3

type Retry struct {
	maxRetries int
}

func NewRetry() *Retry {
	return &Retry{
		maxRetries: 3,
	}
}

func (retry *Retry) Exec(retryFunc func() (*http.Response, bool)) *http.Response {
	var response *http.Response
	var isRetry bool
	for i := 0; i < retry.maxRetries; i++ {
		response, isRetry = retryFunc()
		if !isRetry {
			return response
		}

		log.Println(fmt.Sprintf("retry error. Number retry - %d", i+1))
	}

	return response
}
