package blockchain

import (
	"bytes"
	"encoding/json"

	"github.com/Eric-GreenComb/travelrely-server/bean"
	"github.com/Eric-GreenComb/travelrely-server/config"
	"github.com/Eric-GreenComb/travelrely-server/tools"
)

// PutObj PutObj
func (blockChainAPI *BcAPI) PutObj(fcn, user, org string, peers, args []string) (string, error) {

	var _invokeBean bean.FabricInvokeBean
	for index := range peers {
		_invokeBean.Peers = append(_invokeBean.Peers, peers[index])
	}
	_invokeBean.Fcn = fcn
	for index := range args {
		_invokeBean.Args = append(_invokeBean.Args, args[index])
	}
	_invokeBean.Username = user
	_invokeBean.Orgname = org

	bytesData, err := json.Marshal(_invokeBean)
	if err != nil {
		return "", err
	}

	_resp, err := tools.PostJSONBytes(blockChainAPI.GenPutObjURL(), bytesData)
	if err != nil {
		return "", err
	}
	return _resp, nil
}

// GetObj GetObj
func (blockChainAPI *BcAPI) GetObj(peer, fcn, id, user, org string) (string, error) {
	url := blockChainAPI.GenGetObjURL(peer, fcn, id, user, org)
	_resp, err := tools.HTTPGet(url)
	if err != nil {
		return "", err
	}
	return _resp, nil
}

// GenPutObjURL GenPutObjURL
func (blockChainAPI *BcAPI) GenPutObjURL() string {

	var buffer bytes.Buffer
	buffer.WriteString(config.FabricConfig.APIURL)
	buffer.WriteString("/channels/")
	buffer.WriteString(config.FabricConfig.APIChannel)
	buffer.WriteString("/chaincodes/")
	buffer.WriteString(config.FabricConfig.APIChaincode)

	return buffer.String()
}

// GenGetObjURL GenGetObjURL
func (blockChainAPI *BcAPI) GenGetObjURL(peer, fcn, id, user, org string) string {

	// http://localhost:4000/channels/mychannel1/chaincodes/mycc?peer=peer1&fcn=queryObj&args=%5B%22obj20180101002%22%5D&username=admin&orgname=org1

	var buffer bytes.Buffer
	buffer.WriteString(config.FabricConfig.APIURL)
	buffer.WriteString("/channels/")
	buffer.WriteString(config.FabricConfig.APIChannel)
	buffer.WriteString("/chaincodes/")
	buffer.WriteString(config.FabricConfig.APIChaincode)
	buffer.WriteString("?peer=")
	buffer.WriteString(peer)
	buffer.WriteString("&fcn=")
	buffer.WriteString(fcn)
	buffer.WriteString("&args=%5B%22")
	buffer.WriteString(id)
	buffer.WriteString("%22%5D&username=")
	buffer.WriteString(user)
	buffer.WriteString("&orgname=")
	buffer.WriteString(org)
	return buffer.String()
}
