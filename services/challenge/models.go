package challenge

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type Category struct {
	gorm.Model
	ID         uint           `json:"id" gorm:"primaryKey"`
	UUID       uuid.UUID      `json:"uuid" validate:"required,uuid4"`
	Name       string         `json:"name" validate:"required,string" gorm:"size:127;index:idx_name,unique"`
	CreatedAt  time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"autoCreateTime"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"default:null"`
	Challenges []*Challenge   `gorm:"many2many:category_challenges"`
}

type Challenge struct {
	gorm.Model
	ID          uint           `json:"id"`
	UUID        uuid.UUID      `json:"uuid" validate:"required,uuid4"`
	Name        string         `json:"name" validate:"required,string" gorm:"size:127;index:idx_name,unique"`
	Description string         `json:"description" validate:"required,string" gorm:"size:255;default:null"`
	Street      string         `json:"street" validate:"string" gorm:"size:127;default:null"`
	Postcode    string         `json:"postcode" validate:"string" gorm:"size:10;default:null"`
	Latitude    float32        `json:"latitude" validate:"float" gorm:"default:null"`
	Longitude   float32        `json:"longitude" validate:"float" gorm:"default:null"`
	Thumbnail   string         `json:"image" validate:"string" gorm:"default:null"`
	Gallery     datatypes.JSON `json:"gallery" validate:"slice" gorm:"type:json;default:null"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoCreateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"default:null"`
	Categories  []*Category    `gorm:"many2many:category_challenges"`
}

type CategoryChallenges struct {
	gorm.Model
	CategoryID  uint      `json:"category_id" gorm:"primaryKey"`
	ChallengeID uint      `json:"challenge_id" gorm:"primaryKey"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoCreateTime"`
	DeletedAt   time.Time `json:"deleted_at" gorm:"default:null"`
}
