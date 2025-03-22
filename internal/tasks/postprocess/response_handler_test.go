package postprocess

import (
	"fmt"
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
	s.response.On("GetNeuronRawAnswer").Return("1. **Input:** M = 2, N = 3  \n   - Output: 3  \n2. **Input:** M = 3, N = 2  \n   - Output: 2  \n3. **Input:** M = 4, N = 4  \n   - Output: 4  \n4. **Input:** M = 5, N = 5  \n   - Output: 5  \n5. **Input:** M = 6, N = 6  \n   - Output: 9\",\\\"role\\\":\\\"assistant\\\"},\\\"index\\\":0,\\\"finish_reason\\\":\\\"stop\\\"}],\\\"created\\\":1736018556,\\\"model\\\":\\\"GigaChat:1.0.26.20\\\",\\\"object\\\":\\\"chat.completion\\\",\\\"usage\\\":{\\\"prompt_tokens\\\":113,\\\"completion_tokens\\\":147,\\\"total_tokens\\\":260}}\"")
	s.response.On("AdditionalLogicToHandleResponse").Return(nil)
}

func (s *PostProcessResponseHandlerTestSuite) TestLogic() {
	tests, err := GetTestsByResponse(s.response)
	fmt.Println(tests)
	s.Assert().Nil(err)
}

func TestRun(t *testing.T) {
	suite.Run(t, new(PostProcessResponseHandlerTestSuite))
}
