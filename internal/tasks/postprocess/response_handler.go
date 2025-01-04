package postprocess

import (
	"errors"
	"ginProject1/internal/tasks/config"
	"strings"
)

type TestCase struct {
	input  map[string]string
	output string
}

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

	result := make([]TestCase, count)
	testsRaw := strings.Split(answer, "Входные данные:")[1:]
	for i, testRaw := range testsRaw {
		result[i] = TestCase{}
		inputRaw, outputRaw, ok := strings.Cut(testRaw, "Выходные данные:")
		if !ok {
			return nil, errors.New("Error while parsing test case " + testRaw)
		}
		result[i].input = make(map[string]string)

		for _, el := range strings.Split(strings.TrimSpace(inputRaw), ",") {
			variable, value, ok := strings.Cut(strings.TrimSpace(el), "=")
			if !ok {
				return nil, errors.New("Error while parsing test case " + testRaw)
			}
			newValue, _, ok := strings.Cut(value, "\\")
			if !ok {
				newValue = value
			}
			result[i].input[strings.TrimSpace(variable)] = strings.TrimSpace(newValue)
		}
		newValue, _, ok := strings.Cut(outputRaw, "\\")
		if !ok {
			newValue = outputRaw
		}
		result[i].output = strings.TrimSpace(newValue)
	}

	return result, nil
}
