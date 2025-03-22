package service

import "ginProject1/internal/tasks/repository"

type IListService interface {
	List() (*ModelResponse, error)
}

type ListService struct {
	repository repository.IRepository
}

func NewListService(repository repository.IRepository) *ListService {
	return &ListService{
		repository: repository,
	}
}

func (s *ListService) List() (*ModelResponse, error) {
	models, err := s.repository.ListModels()
	if err != nil {
		return nil, err
	}

	return s.convertToApiModel(models), nil
}

type Model struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func NewModel(id int, name string) Model {
	return Model{
		Id:   id,
		Name: name,
	}
}

type ModelResponse struct {
	Models []Model `json:"data"`
}

func NewModelResponse(models []Model) *ModelResponse {
	return &ModelResponse{
		Models: models,
	}
}

func (s *ListService) convertToApiModel(models []repository.Model) *ModelResponse {
	var answer []Model
	for _, model := range models {
		answer = append(answer, NewModel(model.Id, model.Name))
	}

	return NewModelResponse(answer)
}
