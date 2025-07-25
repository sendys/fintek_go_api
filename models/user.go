package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Uuid     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email" gorm:"unique"`
	Password string    `json:"-"`
}
