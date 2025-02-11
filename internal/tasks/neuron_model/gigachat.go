package neuron_model

import (
	"errors"
	"fmt"
	"ginProject1/internal/tasks/config"
)

var Format = "Входные данные: переменная = значение - Выходные данные: результат. Входные данные разделяй запятой."

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

	gigaChat.Prompt = fmt.Sprintf("Сгенерируй %d наборов тестовых данных для проверки задачи в формате: %s. Текст данной задачи - %s. Ограничения на входные переменные - %s", count, Format, task, conditions)
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
	err := gigaChat.SetModel("GigaChat")
	if err != nil {
		return err
	}
	err = gigaChat.SetAdditionalFields(5000)
	if err != nil {
		return err
	}

	err = gigaChat.GeneratePrompt(task, constraints)
	if err != nil {
		return err
	}
	return nil
}
