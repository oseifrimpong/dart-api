package model

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID        string    `json:"user_id" db:"user_id" gorm:"column:user_id;primary_key;type:uuid"`
	FirstName string    `json:"first_name" db:"first_name" gorm:"varchar(50);column:first_name;not null"`
	LastName  string    `json:"last_name" db:"last_name" gorm:"varchar(50);column:last_name;not null"`
	UserName  string    `json:"user_name" db:"user_name" gorm:"varchar(50);column:user_name;not null;unique"`
	Password  string    `json:"password" db:"password" gorm:"varchar(50);not null"`
	Email     string    `json:"email" db:"email" gorm:"varchar(200);not null;unique"`
	CreatedAt time.Time `json:"created_at" db:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(scope *gorm.Scope) {
	uuid := uuid.NewV4()
	scope.SetColumn("ID", uuid.String())
}
