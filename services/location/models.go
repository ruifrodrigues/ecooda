package location

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Location struct {
	gorm.Model
	ID         uint           `json:"id" gorm:"primaryKey"`
	UUID       uuid.UUID      `json:"uuid" validate:"required,uuid4"`
	ParentID   *uint          `json:"parent_id" gorm:"default:null"`
	Name       string         `json:"name" validate:"required,string" gorm:"size:127;index:idx_name,unique"`
	Type       int32          `json:"type" validate:"required:numeric" gorm:"default:0"`
	CreatedAt  time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"autoCreateTime"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"default:null"`
	Challenges []*LocationChallenges
}

type LocationChallenges struct {
	LocationID    *uint     `json:"location_id" gorm:"primaryKey"`
	ChallengeUUID uuid.UUID `json:"challenge_uuid" gorm:"primaryKey"`
	Location      *Location `gorm:"references:ID"`
}
