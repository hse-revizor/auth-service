package sql

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

var (
	ErrEntityExists   = errors.New("entity already exists")
	ErrEntityNotFound = errors.New("entity not found")
	ErrForeignKey     = errors.New("foreign key error")
)

func isDuplicateError(err error) bool {
	return strings.Contains(err.Error(), "duplicate key value violates unique constraint")
}

func isNotFoundError(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
