package service

import (
	"context"
	"github.com/rajaanova/cleancode/app/entity"
)

// This interface can go in entity package if entity logic demands a call to database
type DataProvider interface {
	GetStudent(ctx context.Context, id string) entity.Student
	AddStudent(ctx context.Context, student entity.Student) error
}