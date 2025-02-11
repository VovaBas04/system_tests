package codechecker

import (
	"errors"
	"fmt"
)

type DtoCodeChecker struct {
	Tests []map[string]string
	Code  string
}

func (dto *DtoCodeChecker) String() string {
	str := ""
	for _, test := range dto.Tests {
		for variable, value := range test {
			str += fmt.Sprintf("%s: %s", variable, value)
		}
	}

	return str
}

func ConvertRequestToDtoCodeCheck(request *TestRequest) (*DtoCodeChecker, error) {
	dtoCodeChecker := &DtoCodeChecker{}
	dtoCodeChecker.Code = request.Code
	dtoCodeChecker.Tests = make([]map[string]string, len(request.Tests))
	for index, test := range request.Tests {
		if _, ok := test["answer"]; !ok {
			return nil, errors.New(fmt.Sprintf("No answer in test - %d", index))
		}

		dtoCodeChecker.Tests[index] = make(map[string]string)
		for variable, value := range test {
			length := len(value)
			if (value[length-1]) != '\n' {
				value = value + "\n"
			}
			dtoCodeChecker.Tests[index][variable] = value
		}
	}

	return dtoCodeChecker, nil
}
