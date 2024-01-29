package service

import (
	"github.com/google/uuid"
	"restfule-api-generic/internal/easywalk/entity"
	"restfule-api-generic/internal/easywalk/repository"
)

type SimplyServiceInterface[T entity.SimplyEntityInterface] interface {
	Create(data T) (*T, error)
	Update(id uuid.UUID, mapFields map[string]any) (int64, error)
	Delete(id uuid.UUID) (int64, error)

	Read(id uuid.UUID) (*T, error)
	FindAll(mapFields map[string]any) ([]T, error)
}

func NewGenericService[T entity.SimplyEntityInterface](repo repository.SimplyRepositoryInterface[T]) SimplyServiceInterface[T] {
	return &simplyService[T]{repo: repo}
}
