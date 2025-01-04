package tasks

import (
	"encoding/json"
	"ginProject1/internal/tasks/postprocess"
	"io"
	"net/http"
)

func ReadNeuronModelAnswer(response *http.Response) (*postprocess.GigaChatResponse, error) {
	defer response.Body.Close()
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	gigaChatResponse := &postprocess.GigaChatResponse{}
	err = json.Unmarshal(bytes, gigaChatResponse)
	if err != nil {
		return nil, err
	}

	return gigaChatResponse, err
}

func ToResponse(testCases []postprocess.TestCase) []map[string]string {
	result := make([]map[string]string, len(testCases))
	for index, testCase := range testCases {
		result[index] = testCase.ToResponse()
	}

	return result
}
