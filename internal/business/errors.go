package business

import (
	"errors"
	"fmt"
)

func NewGetAllItemsError(err error) error {
	return fmt.Errorf("fail to get all items %w", err)
}

func NewGetAnItemsError(err error) error {
	return fmt.Errorf("fail to get an items %w", err)
}

func NewPersonNotFoundError() error {
	return errors.New("person not found")
}
