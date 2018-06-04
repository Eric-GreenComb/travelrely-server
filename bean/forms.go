package bean

import ()

// FormUserRegister FormUserRegister
type FormUserRegister struct {
	CMD    string `form:"cmd" json:"cmd"`
	UserID string `form:"userID" json:"userID"`
}

// FormUserUnRegister FormUserUnRegister
type FormUserUnRegister struct {
	CMD       string `form:"cmd" json:"cmd"`
	BCUserID  string `form:"bcuser_id" json:"bcuser_id"`
	BCUserKey string `form:"bcuser_key" json:"bcuser_key"`
}

// FormSubscribe FormSubscribe
type FormSubscribe struct {
	CMD       string `form:"cmd" json:"cmd"`
	BCUserID  string `form:"bcuser_id" json:"bcuser_id"`
	BCUserKey string `form:"bcuser_key" json:"bcuser_key"`
	Msisdn    string `form:"msisdn" json:"msisdn"`
	Eki2      string `form:"eki2" json:"eki2"`
}

// FormUnsubscribe FormUnsubscribe
type FormUnsubscribe struct {
	CMD       string `form:"cmd" json:"cmd"`
	BCUserID  string `form:"bcuser_id" json:"bcuser_id"`
	BCUserKey string `form:"bcuser_key" json:"bcuser_key"`
	Msisdn    string `form:"msisdn" json:"msisdn"`
	AssetID   string `form:"asset_id" json:"asset_id"`
}

// FormMsisdn FormMsisdn
type FormMsisdn struct {
	CMD    string `form:"cmd" json:"cmd"`
	Msisdn string `form:"msisdn" json:"msisdn"`
}

// FormMsisdns FormMsisdn
type FormMsisdns struct {
	CMD     string   `form:"cmd" json:"cmd"`
	Msisdns []string `form:"msisdn" json:"msisdn"`
}

// FormAsset FormAsset
type FormAsset struct {
	CMD     string `form:"cmd" json:"cmd"`
	AssetID string `form:"asset_id" json:"asset_id"`
}

// FormTransfer FormTransfer
type FormTransfer struct {
	CMD      string `form:"cmd" json:"cmd"`
	BCUserID string `form:"bcuser_id" json:"bcuser_id"`
	Msisdn   string `form:"msisdn" json:"msisdn"`
	AssetID  string `form:"asset_id" json:"asset_id"`
	Op       int8   `form:"op" json:"op"`
}

// FormSearchTx FormSearchTx
type FormSearchTx struct {
	CMD   string `form:"cmd" json:"cmd"`
	Tm    int64  `gorm:"default:0" form:"tm" json:"tm"`
	Start int64  `form:"start" json:"start"`
	MaxNo int64  `form:"max_no" json:"max_no"`
}
