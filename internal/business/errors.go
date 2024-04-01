package business

import (
	"errors"
	"fmt"
)

func NewGetAllItemsError(err error) error {
	return fmt.Errorf("fail to get all items %w", err)
}

func NewGetItemError(err error) error {
	return fmt.Errorf("fail to get an item %w", err)
}

func NewPutItemError(err error) error {
	return fmt.Errorf("fail to put an item %w", err)
}

func NewDeleteItemError(err error) error {
	return fmt.Errorf("fail to delete an item %w", err)
}

func NewPersonNotFoundError() error {
	return errors.New("person not found")
}

func NewFailDecodeError(err error) error {
	return errors.New("fail to decode")
}
