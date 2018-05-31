package blockchain

import (
	"bytes"

	"github.com/Eric-GreenComb/travelrely-server/config"
	"github.com/Eric-GreenComb/travelrely-server/tools"
)

// QueryBlockByBN QueryBlockByBN
func (blockChainAPI *BcAPI) QueryBlockByBN(bn string) (string, error) {

	var buffer bytes.Buffer
	buffer.WriteString(config.FabricConfig.APIURL)
	buffer.WriteString("/channels/")
	buffer.WriteString(config.FabricConfig.APIChannel)
	buffer.WriteString("/blocks/")
	buffer.WriteString(bn)
	buffer.WriteString("?peer=")
	buffer.WriteString(config.FabricConfig.APIAnchor)
	buffer.WriteString("&username=")
	buffer.WriteString("admin")
	buffer.WriteString("&orgname=")
	buffer.WriteString("org1")

	_resp, err := tools.HTTPGet(buffer.String())
	if err != nil {
		return "", err
	}
	return _resp, nil
}

// QueryTrxByID QueryTrxByID
func (blockChainAPI *BcAPI) QueryTrxByID(trxID string) (string, error) {

	var buffer bytes.Buffer
	buffer.WriteString(config.FabricConfig.APIURL)
	buffer.WriteString("/channels/")
	buffer.WriteString(config.FabricConfig.APIChannel)
	buffer.WriteString("/transactions/")
	buffer.WriteString(trxID)
	buffer.WriteString("?peer=")
	buffer.WriteString(config.FabricConfig.APIAnchor)
	buffer.WriteString("&username=")
	buffer.WriteString("admin")
	buffer.WriteString("&orgname=")
	buffer.WriteString("org1")

	_resp, err := tools.HTTPGet(buffer.String())
	if err != nil {
		return "", err
	}
	return _resp, nil
}

// QueryChaincodes QueryChaincodes
func (blockChainAPI *BcAPI) QueryChaincodes(_type string) (string, error) {

	var buffer bytes.Buffer
	buffer.WriteString(config.FabricConfig.APIURL)
	buffer.WriteString("/chaincodes?peer=")
	buffer.WriteString(config.FabricConfig.APIAnchor)
	buffer.WriteString("&username=")
	buffer.WriteString("admin")
	buffer.WriteString("&orgname=")
	buffer.WriteString("org1")
	buffer.WriteString("&type=")
	buffer.WriteString(_type)
	buffer.WriteString("&channel=")
	buffer.WriteString(config.FabricConfig.APIChannel)

	_resp, err := tools.HTTPGet(buffer.String())
	if err != nil {
		return "", err
	}
	return _resp, nil
}

// QueryChannels QueryChannels
func (blockChainAPI *BcAPI) QueryChannels() (string, error) {

	var buffer bytes.Buffer
	buffer.WriteString(config.FabricConfig.APIURL)
	buffer.WriteString("/channels?peer=")
	buffer.WriteString(config.FabricConfig.APIAnchor)
	buffer.WriteString("&username=")
	buffer.WriteString("admin")
	buffer.WriteString("&orgname=")
	buffer.WriteString("org1")

	_resp, err := tools.HTTPGet(buffer.String())
	if err != nil {
		return "", err
	}
	return _resp, nil
}

// QueryChainInfo QueryChainInfo
func (blockChainAPI *BcAPI) QueryChainInfo() (string, error) {

	var buffer bytes.Buffer
	buffer.WriteString(config.FabricConfig.APIURL)
	buffer.WriteString("/channels/")
	buffer.WriteString(config.FabricConfig.APIChannel)
	buffer.WriteString("?peer=")
	buffer.WriteString(config.FabricConfig.APIAnchor)
	buffer.WriteString("&username=")
	buffer.WriteString("admin")
	buffer.WriteString("&orgname=")
	buffer.WriteString("org1")

	_resp, err := tools.HTTPGet(buffer.String())
	if err != nil {
		return "", err
	}
	return _resp, nil
}

// QueryChainHeight QueryChainHeight
func (blockChainAPI *BcAPI) QueryChainHeight(channel string) (string, error) {

	var buffer bytes.Buffer
	buffer.WriteString(config.FabricConfig.APIURL)
	buffer.WriteString("/channels/")
	buffer.WriteString(channel)
	buffer.WriteString("/height?peer=")
	buffer.WriteString(config.FabricConfig.APIAnchor)
	buffer.WriteString("&username=")
	buffer.WriteString("admin")
	buffer.WriteString("&orgname=")
	buffer.WriteString("Org1")

	_resp, err := tools.HTTPGet(buffer.String())
	if err != nil {
		return "", err
	}
	return _resp, nil
}
