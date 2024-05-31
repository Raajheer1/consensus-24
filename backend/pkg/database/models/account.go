package models

import (
	"steller-api/pkg/database"
	"time"
)

type Account struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Email       string    `gorm:"unique;not null" json:"email"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	CompanyName string    `json:"company_name"`
	ProfilePic  string    `json:"profile_pic"`
	PublicKey   string    `json:"public_key"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (a *Account) Create() error {
	return database.DB.Create(a).Error
}

func (a *Account) Get() error {
	if a.PublicKey != "" {
		return database.DB.Where("public_key = ?", a.PublicKey).First(a).Error
	}
	return database.DB.First(a).Error
}

func (a *Account) Update() error {
	return database.DB.Save(a).Error
}

func (a *Account) Delete() error {
	return database.DB.Delete(a).Error
}
