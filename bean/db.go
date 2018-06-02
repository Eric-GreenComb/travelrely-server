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

// Transfers Transfers
type Transfers struct {
	gorm.Model
	Transfer
	Block
}

// Transfer Transfer
type Transfer struct {
	Tm       int64  `gorm:"not null" form:"tm" json:"tm"`
	Msisdn   string `gorm:"default:''" form:"msisdn" json:"msisdn"`
	BcuserID string `gorm:"default:''" form:"bcuser_id" json:"bcuser_id"`
	AssetID  string `gorm:"default:''" form:"asset_id" json:"asset_id"`
	Op       int8   `gorm:"not null" form:"op" json:"op"`
}

// Block Block
type Block struct {
	TxID   string `gorm:"default:''" form:"txid" json:"txid"`
	Height int64  `gorm:"not null" form:"height" json:"height"`
}
