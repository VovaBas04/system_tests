package codechecker

type TestRequest struct {
	Tests []map[string]string `json:"tests" binding:"required,dive"`
	Code  string              `json:"code" binding:"required"`
}

func NewTestRequest() *TestRequest {
	return &TestRequest{}
}
