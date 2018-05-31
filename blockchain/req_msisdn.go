package blockchain

import (
	"github.com/Eric-GreenComb/travelrely-server/config"
)

// SubscribeMsisdn SubscribeMsisdn
func (blockChainAPI *BcAPI) SubscribeMsisdn(msisdn, assetID, eki2, bcuserID, bcuserKey string) (string, error) {

	var _peers []string
	_peers = append(_peers, config.FabricConfig.APIAnchor)

	var _args []string
	_args = append(_args, msisdn)
	_args = append(_args, assetID)
	_args = append(_args, eki2)
	_args = append(_args, bcuserID)
	_args = append(_args, bcuserKey)

	return blockChainAPI.PutObj("subscribe", "admin", "Org1", _peers, _args)
}

// UnsubscribeMsisdn UnsubscribeMsisdn
func (blockChainAPI *BcAPI) UnsubscribeMsisdn(msisdn, assetID, bcuseID, bcuserKey string) (string, error) {

	var _peers []string
	_peers = append(_peers, config.FabricConfig.APIAnchor)

	var _args []string
	_args = append(_args, msisdn)
	_args = append(_args, assetID)
	_args = append(_args, bcuseID)
	_args = append(_args, bcuserKey)

	return blockChainAPI.PutObj("unsubscribe", "admin", "Org1", _peers, _args)
}

// GetMsisdnState GetMsisdnState
func (blockChainAPI *BcAPI) GetMsisdnState(msisdn string) (string, error) {
	return blockChainAPI.GetObj(config.FabricConfig.APIAnchor, "msisdn_state", msisdn, "admin", "Org1")
}

// GetMsisdnHistory GetMsisdnHistory
func (blockChainAPI *BcAPI) GetMsisdnHistory(msisdn string) (string, error) {
	return blockChainAPI.GetObj(config.FabricConfig.APIAnchor, "get_msisdn_history", msisdn, "admin", "Org1")
}

// GetAssetInfo GetAssetInfo
func (blockChainAPI *BcAPI) GetAssetInfo(asset string) (string, error) {
	return blockChainAPI.GetObj(config.FabricConfig.APIAnchor, "asset_info", asset, "admin", "Org1")
}
