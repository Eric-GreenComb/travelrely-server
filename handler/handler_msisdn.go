package handler

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/travelrely-server/bean"
	"github.com/Eric-GreenComb/travelrely-server/blockchain"
	"github.com/Eric-GreenComb/travelrely-server/config"
	"github.com/Eric-GreenComb/travelrely-server/persist"
	"github.com/Eric-GreenComb/travelrely-server/tools"
)

// SubscribeMsisdn SubscribeMsisdn
func SubscribeMsisdn(c *gin.Context) {

	_formSubscribe := bean.FormSubscribe{}
	c.BindJSON(&_formSubscribe)

	_key, err := tools.Ks.NewKey()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 11, "msg": err.Error()})
		return
	}

	_asset := _key.Address.String()
	_resp, err := blockchain.GetBCAPI().SubscribeMsisdn(_formSubscribe.Msisdn, _asset, _formSubscribe.Eki2, _formSubscribe.BCUserID, _formSubscribe.BCUserKey)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "msg": err.Error()})
	}

	if len(_resp) == 0 {
		c.JSON(200, gin.H{"errcode": 1, "msg": "response is empty"})
		return
	}
	if strings.Contains(_resp, "cc_error") {
		c.JSON(200, gin.H{"errcode": 1, "msg": "cc_error"})
		return
	}

	_height, err := blockchain.GetBCAPI().QueryChainHeight(config.FabricConfig.APIChannel)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 0, "txid": _resp, "asset": _asset})
		return
	}

	_iHeight, _ := strconv.ParseInt(_height, 10, 64)

	var _tx bean.Transfers

	_tx.Tm = time.Now().Unix()
	_tx.Msisdn = _formSubscribe.Msisdn
	_tx.BcuserID = _formSubscribe.BCUserID
	_tx.AssetID = _asset
	_tx.Op = 0
	_tx.Height = _iHeight
	_tx.TxID = _resp

	err = persist.GetPersist().CreateTransfer(_tx)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 0, "txid": _resp, "asset": _asset})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "txid": _resp, "asset": _asset})
}

// UnsubscribeMsisdn UnsubscribeMsisdn
func UnsubscribeMsisdn(c *gin.Context) {

	_formUnsubscribe := bean.FormUnsubscribe{}
	c.BindJSON(&_formUnsubscribe)

	_resp, err := blockchain.GetBCAPI().UnsubscribeMsisdn(_formUnsubscribe.Msisdn, _formUnsubscribe.AssetID, _formUnsubscribe.BCUserID, _formUnsubscribe.BCUserKey)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "msg": err.Error()})
	}

	if len(_resp) == 0 {
		c.JSON(200, gin.H{"errcode": 1, "msg": "response is empty"})
		return
	}
	if strings.Contains(_resp, "cc_error") {
		c.JSON(200, gin.H{"errcode": 1, "msg": "cc_error"})
		return
	}

	_height, err := blockchain.GetBCAPI().QueryChainHeight(config.FabricConfig.APIChannel)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 0, "txid": _resp})
		return
	}

	_iHeight, _ := strconv.ParseInt(_height, 10, 64)

	var _tx bean.Transfers

	_tx.Tm = time.Now().Unix()
	_tx.Msisdn = _formUnsubscribe.Msisdn
	_tx.BcuserID = _formUnsubscribe.BCUserID
	_tx.AssetID = _formUnsubscribe.AssetID
	_tx.Op = 1
	_tx.Height = _iHeight
	_tx.TxID = _resp

	err = persist.GetPersist().CreateTransfer(_tx)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 0, "txid": _resp})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "txid": _resp})
}

// GetMsisdnState GetMsisdnState
func GetMsisdnState(c *gin.Context) {

	_msisdn := c.Params.ByName("msisdn")

	_resp, err := blockchain.GetBCAPI().GetMsisdnState(_msisdn)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "msg": err.Error()})
	}

	if len(_resp) == 0 {
		c.JSON(200, gin.H{"errcode": 1, "msg": "response is empty"})
	} else {
		c.String(http.StatusOK, _resp)
	}
}

// GetMsisdnHistory GetMsisdnHistory
func GetMsisdnHistory(c *gin.Context) {

	_msisdn := c.Params.ByName("msisdn")

	_resp, err := blockchain.GetBCAPI().GetMsisdnHistory(_msisdn)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "msg": err.Error()})
	}

	if len(_resp) == 0 {
		c.JSON(200, gin.H{"errcode": 1, "msg": "response is empty"})
	} else {
		c.String(http.StatusOK, _resp)
	}
}

// GetAssetInfo GetAssetInfo
func GetAssetInfo(c *gin.Context) {

	_asset := c.Params.ByName("asset")

	_resp, err := blockchain.GetBCAPI().GetAssetInfo(_asset)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "msg": err.Error()})
	}

	if len(_resp) == 0 {
		c.JSON(200, gin.H{"errcode": 1, "msg": "response is empty"})
	} else {
		c.String(http.StatusOK, _resp)
	}
}
