package blockchain

import (
	"github.com/Eric-GreenComb/travelrely-server/config"
)

// CreateObjByID CreateObjByID
func (blockChainAPI *BcAPI) CreateObjByID(ID, base64 string) (string, error) {

	var _peers []string
	_peers = append(_peers, "localhost:7051")

	var _args []string
	_args = append(_args, ID)
	_args = append(_args, base64)

	return blockChainAPI.PutObj("createObj", "admin", "org1", _peers, _args)
}

// UpdateObjByID UpdateObjByID
func (blockChainAPI *BcAPI) UpdateObjByID(ID, base64 string) (string, error) {

	var _peers []string
	_peers = append(_peers, "localhost:7051")

	var _args []string
	_args = append(_args, ID)
	_args = append(_args, base64)

	return blockChainAPI.PutObj("updateObj", "admin", "org1", _peers, _args)
}

// GetObjByID GetObjByID
func (blockChainAPI *BcAPI) GetObjByID(ID string) (string, error) {
	return blockChainAPI.GetObj(config.FabricConfig.APIAnchor, "queryObj", ID, "admin", "org1")
}

// GetObjHistoryByID GetObjHistoryByID
func (blockChainAPI *BcAPI) GetObjHistoryByID(ID string) (string, error) {
	return blockChainAPI.GetObj(config.FabricConfig.APIAnchor, "getObjHistory", ID, "admin", "org1")
}
