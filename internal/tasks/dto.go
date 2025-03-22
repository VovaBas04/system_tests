package tasks

import (
	"ginProject1/internal/tasks/neuron_model"
	"log"
)

type PromptDto struct {
	task        string
	model       string
	constraints map[string]map[neuron_model.Constraint]string
}

func ConvertRequestToDtoConstraint(request *PromptRequest) *PromptDto {
	promptDto := &PromptDto{}
	promptDto.task = request.Prompt
	promptDto.model = request.Model
	result := make(map[string]map[neuron_model.Constraint]string)
	for _, constraints := range request.Params {
		if _, ok := result[constraints.Variable]; !ok {
			result[constraints.Variable] = make(map[neuron_model.Constraint]string)
		}
		log.Print(constraints)
		result[constraints.Variable][constraints.Constraint] = constraints.Value
	}
	promptDto.constraints = result

	return promptDto
}
