package postprocess

import (
	"errors"
	"fmt"
)

type TestCase struct {
	input  map[string]string
	output string
}

type ParseTest interface {
	Parse(answer string, count int) ([]TestCase, error)
}

type Handler struct {
	handlers []ParseTest
}

func NewHandler() *Handler {
	return &Handler{
		handlers: []ParseTest{
			NewMainStrategy(),
			NewNumerateStrategy(),
		},
	}
}

func (h *Handler) Handle(answer string, count int) ([]TestCase, error) {
	for _, handler := range h.handlers {
		testCases, err := handler.Parse(answer, count)
		if err == nil {
			return testCases, nil
		}
		fmt.Print(err)
	}
	return nil, errors.New("not parsed answer")
}
