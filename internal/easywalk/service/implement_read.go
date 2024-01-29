package service

import "github.com/google/uuid"

func (s simplyService[T]) Read(id uuid.UUID) (*T, error) {
	data, err := s.repo.Read(id)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
