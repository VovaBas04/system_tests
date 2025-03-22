package postprocess

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type NumerateStrategy struct{}

func NewNumerateStrategy() *NumerateStrategy {
	return &NumerateStrategy{}
}

func (h *NumerateStrategy) Parse(answer string, count int) ([]TestCase, error) {
	regex, err := regexp.Compile(fmt.Sprintf("[1-%d][.)] ?", count))
	if err != nil {
		return nil, err
	}
	testCasesRaw := regex.Split(answer, count+1)[1:]
	lenTests := len(testCasesRaw)
	if lenTests < 5 {
		return nil, errors.New("number of tests must be at least 5")
	}
	
	if lenTests > 5 {
		testCasesRaw = testCasesRaw[:5]
	}

	result := make([]TestCase, count)
	for i, testRaw := range testCasesRaw {
		result[i] = TestCase{}
		inputRaw, outputRaw, ok := strings.Cut(testRaw, "-")
		if !ok {
			return nil, errors.New("Error while parsing test case " + testRaw)
		}
		result[i].input = make(map[string]string)

		for _, el := range strings.Split(strings.Trim(strings.TrimSpace(inputRaw), "-"), ",") {
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
		regexpOutput, err := regexp.Compile("Output:? ?")
		if err != nil {
			return nil, err
		}

		outputRaw = regexpOutput.ReplaceAllString(outputRaw, "")
		newValue, _, ok := strings.Cut(outputRaw, "\\")
		if !ok {
			newValue = outputRaw
		}
		result[i].output = strings.TrimSpace(newValue)
	}

	return result, nil
}
