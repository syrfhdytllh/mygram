package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name             string `gorm:"varchar(255);not null" json:"name" valid:"required~Your name is required"`
	Social_media_url string `gorm:"varchar(255);not null" json:"socialMediaUrl" valid:"required~Your socialMediaUrl is required"`
	UserId           uint
	User             *User
	// CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (p *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return

}
