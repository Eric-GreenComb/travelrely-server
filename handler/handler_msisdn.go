package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/travelrely-server/bean"
	"github.com/Eric-GreenComb/travelrely-server/blockchain"
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
