package user

import "github.com/google/uuid"

type User struct {
	Id          uuid.UUID `gorm:"column:id;primary_key" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	Email       string    `gorm:"column:email" json:"email"`
	PhoneNumber string    `gorm:"column:phone_number" json:"phoneNumber"`
	Address     string    `gorm:"column:address" json:"address"`
}
