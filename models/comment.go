package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	// User_id  []User  `gorm:"foreignKey:Id" json:"userId"`
	// Photo_id []Photo `gorm:"foreignKey:Id" json:"photoId"`
	UserId  uint
	PhotoId uint   `json:"photo_id" form:"photo_id"`
	Message string `gorm:"varchar(255);not null" json:"message" valid:"required~Your message is required"`
	User    *User
	Photo   *Photo
	// CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (p *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return

}
