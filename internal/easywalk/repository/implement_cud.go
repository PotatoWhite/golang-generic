package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"reflect"
	"restfule-api-generic/internal/easywalk/entity"
)

type simplyRepository[T entity.SimplyEntityInterface] struct {
	db *gorm.DB
}

// Create is a generic method for create operation
// @param data - pointer to entity
// @return uuid of created entity, error
func (r *simplyRepository[T]) Create(data T) (uuid.UUID, error) {
	data.SetID(uuid.New())
	tx := r.db.Create(&data)
	if tx.Error != nil {
		log.Printf("Error in repository Create operation - %v", tx.Error)
	}

	return data.GetID(), tx.Error
}

func (r *simplyRepository[T]) Update(id uuid.UUID, mapFields map[string]any) (int64, error) {

	fromDB, err := r.Read(id)
	if err != nil {
		return 0, err
	}

	// print all fields of T
	for i := 0; i < reflect.TypeOf(fromDB).NumField(); i++ {
		field := reflect.TypeOf(fromDB).Field(i)
		log.Printf("Field: %s, Type: %s", field.Name, field.Type)
	}

	// set all fields of T from mapFields
	for key, value := range mapFields {
		reflect.ValueOf(&fromDB).Elem().FieldByName(key).Set(reflect.ValueOf(value))
	}
	// update T
	tx := r.db.Save(&fromDB)
	if tx.Error != nil {
		log.Printf("Error in repository Update operation - %v", tx.Error)
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (r *simplyRepository[T]) Delete(id uuid.UUID) (int64, error) {
	var deleted T
	tx := r.db.Delete(&deleted, id)
	if tx.Error != nil {
		log.Printf("Error in repository Delete operation - %v", tx.Error)
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}
