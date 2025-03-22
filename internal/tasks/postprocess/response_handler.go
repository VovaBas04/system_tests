package postprocess

import (
	"ginProject1/internal/tasks/config"
)

func (testCase *TestCase) ToResponse() map[string]string {
	result := make(map[string]string)
	for key, value := range testCase.input {
		result[key] = value
	}
	result["answer"] = testCase.output

	return result
}

func GetTestsByResponse(response NeuronModelResponse) ([]TestCase, error) {
	err := response.AdditionalLogicToHandleResponse()
	if err != nil {
		return nil, err
	}

	answer := response.GetNeuronRawAnswer()
	count, err := config.GetTestCases()
	if err != nil {
		return nil, err
	}

	handler := NewHandler()

	return handler.Handle(answer, count)
}
