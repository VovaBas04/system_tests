package postprocess

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type MainStrategy struct{}

func NewMainStrategy() *MainStrategy {
	return &MainStrategy{}
}

func (m *MainStrategy) Parse(answer string, count int) ([]TestCase, error) {
	result := make([]TestCase, count)
	re, err := regexp.Compile(fmt.Sprintf("[1-%d][.)-] ?", count))
	if err != nil {
		return nil, err
	}
	answerWithoutNumerate := re.ReplaceAllString(answer, "")
	testsRaw := strings.Split(strings.ReplaceAll(strings.ReplaceAll(answerWithoutNumerate, "*", ""), "Objective:", ""), "Input:")[1:]
	lenTests := len(testsRaw)
	if lenTests < 5 {
		return nil, errors.New("number of tests must be at least 5")
	}
	if lenTests > 5 {
		testsRaw = testsRaw[:5]
	}

	for i, testRaw := range testsRaw {
		result[i] = TestCase{}
		inputRaw, outputRaw, ok := strings.Cut(testRaw, "Output:")
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
		newValue, _, ok := strings.Cut(outputRaw, "\\")
		if !ok {
			newValue = outputRaw
		}
		result[i].output = strings.TrimSpace(newValue)
	}
	err = checkAnswer(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func checkAnswer(testCases []TestCase) error {
	for ind, testCase := range testCases {
		if testCase.output == "" {
			return errors.New(fmt.Sprintf("empty answer - %d", ind))
		}
	}

	return nil
}
