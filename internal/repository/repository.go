package repository

import "github.com/lucaspereirasilva0/rest-api/internal/model"

type Service interface {
	GetAllItems()([]model.Person, error)
}