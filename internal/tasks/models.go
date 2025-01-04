package tasks

import "ginProject1/internal/tasks/neuron_model"

type PromptRequest struct {
	Prompt string `binding:"required" json:"prompt"`
	Params []struct {
		Variable   string                  `binding:"required" json:"variable"`
		Value      string                  `binding:"required" json:"value"`
		Constraint neuron_model.Constraint `binding:"required" json:"constraint"`
	} `binding:"dive" json:"params"`
}
