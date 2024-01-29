package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID   uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name string    `gorm:"type:varchar(255);not null;unique;"`
	Age  int
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) GetID() uuid.UUID {
	return u.ID
}

func (u *User) SetID(id uuid.UUID) {
	u.ID = id
}
