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
