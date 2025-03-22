package neuron_model

import (
	"errors"
	"fmt"
	"ginProject1/internal/tasks/config"
)

var Format = "\"Input: переменная1 = значение1, переменная2 = значение2, ..., переменнаяN = значениеN - Output: результат\". Input данные разделяй запятой. Input от Output разделяй тире. Нужно точное соответствие заданному формату без лишней информации!"

type GigaChatRequest struct {
	Model     string
	MaxTokens int
	Prompt    string
}

func NewGigaChatRequest() *GigaChatRequest {
	return &GigaChatRequest{}
}

func (gigaChat *GigaChatRequest) GeneratePrompt(task string, constraints map[string]map[Constraint]string) error {
	var conditions string
	for variable, constraint := range constraints {
		for operator, value := range constraint {
			conditions += fmt.Sprintf("%s %s %s, ", variable, operator, value)
		}
	}
	count, err := config.GetTestCases()
	if err != nil {
		return err
	}
	if conditions != "" {
		conditions = " Ограничения на входные переменные -" + conditions
	}

	gigaChat.Prompt = fmt.Sprintf("Сгенерируй наборы тестовых данных для проверки задачи в формате: %s.\"%s\". %s. Количество тестов должны быть равным - %d", Format, task, conditions, count)

	return nil
}

func (gigaChat *GigaChatRequest) SetModel(model string) error {
	gigaChat.Model = model

	return nil
}

func (gigaChat *GigaChatRequest) SetAdditionalFields(args ...interface{}) error {
	if _, ok := args[0].(int); !ok || len(args) > 1 {
		return errors.New("neuron gigachat requires an integer argument")
	}
	gigaChat.MaxTokens = args[0].(int)

	return nil
}

func (gigaChat *GigaChatRequest) PreparedRequest(task string, constraints map[string]map[Constraint]string) error {
	err := gigaChat.SetAdditionalFields(5000)
	if err != nil {
		return err
	}

	err = gigaChat.GeneratePrompt(task, constraints)
	if err != nil {
		return err
	}
	
	return nil
}
