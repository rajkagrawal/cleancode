package service

import (
	"context"
	"github.com/rajaanova/cleancode/app/entity"
)



type Service struct {
	db DataProvider
}
func NewService(db DataProvider) *Service {
	return &Service{db}
}
func (a *Service) CreateStudent(st entity.Student) error {
	ctx := context.Background()
	a.db.AddStudent(ctx, st)
	return nil
}

