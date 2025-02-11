package postprocess

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type PostProcessResponseHandlerTestSuite struct {
	suite.Suite
	response *MockResponse
}

type MockResponse struct {
	mock.Mock
}

func (mock *MockResponse) GetNeuronRawAnswer() string {
	args := mock.Called()

	return args.String(0)
}

func (mock *MockResponse) AdditionalLogicToHandleResponse() error {
	args := mock.Called()

	return args.Error(0)
}

func (s *PostProcessResponseHandlerTestSuite) SetupTest() {
	err := os.Chdir("../../../")
	s.Assert().Nil(err)
	s.response = &MockResponse{}
	s.response.On("GetNeuronRawAnswer").Return("\"Входные данные: n = 5, k = 3\\\\n- Выходные данные: 10\\\\n\\\\nВходные данные: n = 10, k = 5\\\\n- Выходные данные: 252\\\\n\\\\nВходные данные: n = 50, k = 25\\\\n- Выходные данные: 1225\\\\n\\\\nВходные данные: n = 100, k = 50\\\\n- Выходные данные: 126 \\\\n\\\\nВходные данные: n = 200, k = 100\\\\n- Выходные данные: 19075\\\",\\\"role\\\":\\\"assistant\\\"},\\\"index\\\":0,\\\"finish_reason\\\":\\\"stop\\\"}],\\\"created\\\":1736018556,\\\"model\\\":\\\"GigaChat:1.0.26.20\\\",\\\"object\\\":\\\"chat.completion\\\",\\\"usage\\\":{\\\"prompt_tokens\\\":113,\\\"completion_tokens\\\":147,\\\"total_tokens\\\":260}}\"")
	s.response.On("AdditionalLogicToHandleResponse").Return(nil)
}

func (s *PostProcessResponseHandlerTestSuite) TestLogic() {
	_, err := GetTestsByResponse(s.response)
	s.Assert().Nil(err)
}

func TestRun(t *testing.T) {
	suite.Run(t, new(PostProcessResponseHandlerTestSuite))
}
