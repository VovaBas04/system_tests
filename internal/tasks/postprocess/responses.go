package postprocess

type NeuronModelResponse interface {
	GetNeuronRawAnswer() string
	AdditionalLogicToHandleResponse() error
}
type GigaChatResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
			Role    string `json:"role"`
		} `json:"message"`
		Index        int    `json:"index"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Object  string `json:"object"`
	Usage   struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

func (response *GigaChatResponse) GetNeuronRawAnswer() string {
	return response.Choices[0].Message.Content
}

func (response *GigaChatResponse) AdditionalLogicToHandleResponse() error {
	return nil
}
