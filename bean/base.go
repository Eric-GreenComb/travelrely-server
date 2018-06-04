package bean

import ()

// FabricInvokeBean FabricInvokeBean
type FabricInvokeBean struct {
	Peers    []string `json:"peers"`    // 节点组
	Fcn      string   `json:"fcn"`      // 调用函数
	Args     []string `json:"args"`     // 参数
	Username string   `json:"username"` // 组织用户
	Orgname  string   `json:"orgname"`  // 组织
}

// TxHistory TxHistory
type TxHistory struct {
	TxID      string `json:"txid,omitempty"`
	Value     string `json:"value,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	IsDelete  string `json:"isdelete,omitempty"`
}

// TxHistoryOut TxHistoryOut
type TxHistoryOut struct {
	TxID      string `json:"txid,omitempty"`
	Data      Msisdn `json:"data,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	IsDelete  string `json:"isdelete,omitempty"`
}

// Msisdn Msisdn
type Msisdn struct {
	Msisdn  string `json:"msisdn,omitempty"`
	AssetID string `json:"asset_id,omitempty"`
	UserID  string `json:"user_id,omitempty"`
	UserKey string `json:"user_key,omitempty"`
	Status  int    `json:"status,omitempty"`
}

// MsisdnStates MsisdnStates
type MsisdnStates struct {
	Msisdn  string `json:"msisdn,omitempty"`
	AssetID string `json:"asset_id,omitempty"`
	Status  int    `json:"status,omitempty"`
}
