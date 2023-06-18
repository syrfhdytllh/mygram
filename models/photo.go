package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title     string `gorm:"varchar(255);not null" json:"title" valid:"required~Your title is required"`
	Caption   string `gorm:"varchar(255)" json:"caption"`
	Photo_url string `gorm:"varchar(255);not null" json:"photoUrl" valid:"required~Your photoUrl is required"`
	UserId    uint
	User      *User
	// CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return

}
