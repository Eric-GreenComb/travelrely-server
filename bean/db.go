package bean

import (
	"github.com/jinzhu/gorm"
)

// Users Users
type Users struct {
	gorm.Model
	User
}

// User User
type User struct {
	UserID     string `gorm:"default:''" form:"userID" json:"userID"`
	Address    string `gorm:"default:''" form:"address" json:"address"`              // 用户地址
	PrivateKey string `gorm:"size:600;not null" form:"PrivateKey" json:"PrivateKey"` // 用户PrivateKey
}
