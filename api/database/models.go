package database

import (
	"github.com/google/uuid"
)

type Clothing struct {
	Id       uuid.UUID `json:"id" gorm:"primary_key; type:uuid; default:gen_random_uuid()"`
	Name     string    `json:"name" gorm:"not null; varchar(100)"`
	Color    string    `json:"color" gorm:"not null; varchar(100)"`
	Category string    `json:"category" gorm:"not null; varchar(100)"`
	Image    string    `json:"image" gorm:"not null; varchar(255)"`
}

type Outfit struct {
	Id       uuid.UUID  `json:"id" gorm:"primary_key; type:uuid; default:gen_random_uuid()"`
	Name     string     `json:"name" gorm:"not null; varchar(100)"`
	Clothing []Clothing `json:"clothing" gorm:"many2many:outfit_clothing;"`
}
