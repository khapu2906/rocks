package entities

import (
	"time"
	"github.com/lib/pq"
)

type User struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Username       string    `gorm:"unique;not null" json:"username"`
	Email          string    `gorm:"unique;not null" json:"email"`
	EmailVerified  bool      `gorm:"default:false" json:"email_verified"`
	Password       string    `gorm:"not null" json:"-"`
	FullName       string    `json:"full_name,omitempty"`
	PhoneNumber    string    `gorm:"type:varchar(20)" json:"phone_number,omitempty"`
	AvatarURL      string    `json:"avatar_url,omitempty"`
	Role           string    `gorm:"type:varchar(50);default:'user'" json:"role"`
	Status         string    `gorm:"type:varchar(20);default:'active'" json:"status"`
	LastLoginAt    *time.Time `json:"last_login_at,omitempty"`
	AuthProvider   string    `gorm:"type:varchar(50);default:'local'" json:"auth_provider"`

	Metadata       JSONB     `gorm:"type:jsonb;default:'{}'" json:"metadata,omitempty"`

	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `gorm:"index" json:"-"`
}
