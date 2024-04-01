package business

import (
	"log"

	"github.com/lucaspereirasilva0/rest-api/internal/model"
	"github.com/lucaspereirasilva0/rest-api/internal/repository"
)

type service struct {
	repository repository.Service
}

func NewService(repository repository.Service) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetPerson() ([]model.Person, error) {
	personList, err := s.repository.GetAllItems()
	if err != nil {
		log.Println(err)
		return []model.Person{}, NewGetAllItemsError(err)
	}
	return personList, nil
}
