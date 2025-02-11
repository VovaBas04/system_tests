package service

import "ginProject1/internal/tasks/repository"

type IListService interface {
	List() ([]Model, error)
}

type ListService struct {
	repository repository.IRepository
}

func NewListService(repository repository.IRepository) *ListService {
	return &ListService{
		repository: repository,
	}
}

func (s *ListService) List() ([]Model, error) {
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

func (s *ListService) convertToApiModel(models []repository.Model) []Model {
	var answer []Model
	for _, model := range models {
		answer = append(answer, NewModel(model.Id, model.Name))
	}

	return answer
}
