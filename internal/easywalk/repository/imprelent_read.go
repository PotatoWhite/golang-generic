package repository

import (
	"github.com/google/uuid"
	"log"
)

func (r *simplyRepository[T]) Read(id uuid.UUID) (T, error) {
	var out T
	tx := r.db.First(&out, id)
	return out, tx.Error
}

func (r *simplyRepository[T]) ReadAll() ([]T, error) {
	var out []T
	tx := r.db.Find(&out)
	return out, tx.Error
}

func (r *simplyRepository[T]) FindAll(mapFields map[string]any) ([]T, error) {
	var out []T
	tx := r.db.Where(mapFields).Find(&out)
	if tx.Error != nil {
		log.Printf("Error in repository FindAll operation - %v", tx.Error)
	}
	return out, tx.Error
}
