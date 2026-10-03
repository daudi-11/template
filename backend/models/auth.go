package models

import "time"

// AdminUser is retained so existing primary accounts can continue to sign in.
type AdminUser struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	Email        string    `json:"email" gorm:"uniqueIndex;not null"`
	PasswordHash string    `json:"-" gorm:"not null"`
	IsPrimary    bool      `json:"is_primary" gorm:"not null;default:false;index"`
	CreatedAt    time.Time `json:"created_at"`
}

type AuthCredential struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Name         string `json:"name"`
	Email        string `json:"email" gorm:"uniqueIndex;not null"`
	PasswordHash string `json:"-" gorm:"not null"`
	// AccountType keeps the existing table compatible; new template accounts use the generic value "user".
	AccountType string `json:"-" gorm:"index;not null;default:user"`
	CreatedAt   time.Time
}

type SessionToken struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	TokenHash string     `json:"-" gorm:"uniqueIndex;not null"`
	UserEmail string     `json:"user_email" gorm:"index;not null"`
	ExpiresAt time.Time  `json:"expires_at" gorm:"index;not null"`
	RevokedAt *time.Time `json:"revoked_at"`
	CreatedAt time.Time
}
