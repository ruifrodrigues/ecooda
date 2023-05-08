package challenge

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Challenge struct {
	gorm.Model
	ID          uint           `json:"id"`
	LocationID  uint           `json:"locationId" validate:"required,number"`
	CategoryID  uint           `json:"categoryId" validate:"required,number"`
	Name        string         `json:"name" validate:"required,string" gorm:"size:255;index:idx_name,unique"`
	Description string         `json:"description" validate:"required,string"`
	Image       string         `json:"image" validate:"string" gorm:"default:null"`
	Gallery     datatypes.JSON `json:"gallery" validate:"slice" gorm:"type:json;default:null"`
}
