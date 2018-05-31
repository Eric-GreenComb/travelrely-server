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
