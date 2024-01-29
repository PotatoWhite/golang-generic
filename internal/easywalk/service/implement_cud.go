package service

import (
	"github.com/google/uuid"
	"restfule-api-generic/internal/easywalk/entity"
	"restfule-api-generic/internal/easywalk/repository"
)

type simplyService[T entity.SimplyEntityInterface] struct {
	repo repository.SimplyRepositoryInterface[T]
}

func (s simplyService[T]) Create(data T) (*T, error) {
	id, err := s.repo.Create(data)
	if err != nil {
		return nil, err
	}

	// get data from db
	data, err = s.repo.Read(id)
	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (s simplyService[T]) Update(id uuid.UUID, mapFields map[string]any) (int64, error) {
	affected, err := s.repo.Update(id, mapFields)
	if err != nil {
		return 0, err
	}

	return affected, nil
}

func (s simplyService[T]) Delete(id uuid.UUID) (int64, error) {
	affected, err := s.repo.Delete(id)
	if err != nil {
		return 0, err
	}

	return affected, nil
}

func (s simplyService[T]) FindAll(mapFields map[string]any) ([]T, error) {
	data, err := s.repo.FindAll(mapFields)
	if err != nil {
		return nil, err
	}

	return data, nil
}
