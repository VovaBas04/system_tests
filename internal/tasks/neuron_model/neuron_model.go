package neuron_model

import (
	"net/http"
)

type NeuronModelRequest interface {
	GeneratePrompt(task string, constraints map[string]map[Constraint]string) error
	SetModel(model string) error
	SetAdditionalFields(args ...interface{}) error
	PreparedRequest() error
	Do() (*http.Response, error)
}
