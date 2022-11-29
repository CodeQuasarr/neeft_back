package users

/**
 * @Author ANYARONKE Daré Samuel
 */

import "time"

type User struct {
	ID              uint   `gorm:"primaryKey"   json:"id" `
	Username        string `gorm:"varchar(255)" json:"username"`
	FirstName       string `gorm:"varchar(255)" json:"first_name"`
	LastName        string `gorm:"varchar(255)" json:"last_name"`
	Email           string `gorm:"varchar(255)" json:"email"`
	EmailVerifiedAt bool   `gorm:"boolean"      json:"email_verified_at"`
	Password        string `gorm:"varchar(255)" json:"password"`
	Created_at      time.Time
	Updated_at      time.Time
	Deleted_at      time.Time
}
