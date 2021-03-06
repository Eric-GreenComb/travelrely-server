package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	mystrings "github.com/Eric-GreenComb/contrib/strings"
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

	fmt.Println("===== SubscribeMsisdn =====")
	fmt.Println(time.Now())
	fmt.Println(_formSubscribe)
	fmt.Println("===== SubscribeMsisdn end =====")

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
		c.JSON(200, gin.H{"errcode": 1, "msg": _resp})
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

	fmt.Println("===== SubscribeMsisdn =====")
	fmt.Println(time.Now())
	fmt.Println(_formUnsubscribe)
	fmt.Println("===== SubscribeMsisdn end =====")

	_resp, err := blockchain.GetBCAPI().UnsubscribeMsisdn(_formUnsubscribe.Msisdn, _formUnsubscribe.AssetID, _formUnsubscribe.BCUserID, _formUnsubscribe.BCUserKey)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "msg": err.Error()})
	}

	if len(_resp) == 0 {
		c.JSON(200, gin.H{"errcode": 1, "msg": "response is empty"})
		return
	}
	if strings.Contains(_resp, "cc_error") {
		c.JSON(200, gin.H{"errcode": 1, "msg": _resp})
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

	_formMsisdn := bean.FormMsisdn{}
	c.BindJSON(&_formMsisdn)

	_encode := url.QueryEscape(_formMsisdn.Msisdn)

	_resp, err := blockchain.GetBCAPI().GetMsisdnState(_encode)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "msg": err.Error()})
	}

	if len(_resp) == 0 {
		c.JSON(200, gin.H{"errcode": 1, "msg": "response is empty"})
	} else {
		c.String(http.StatusOK, _resp)
	}
}

// GetMsisdnStates GetMsisdnStates
func GetMsisdnStates(c *gin.Context) {

	_formMsisdns := bean.FormMsisdns{}
	c.BindJSON(&_formMsisdns)

	var _msisdns string
	for _, _msisdn := range _formMsisdns.Msisdns {
		_encode := url.QueryEscape(_msisdn)
		_msisdns = _msisdns + _encode + ","
	}

	if len(_msisdns) > 0 {
		_msisdns = mystrings.Substr(_msisdns, 0, len(_msisdns)-1)
	}

	_resp, err := blockchain.GetBCAPI().GetMsisdnStates(_msisdns)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "msg": err.Error()})
	}

	if len(_resp) == 0 {
		c.JSON(200, gin.H{"errcode": 1, "msg": "response is empty"})
		return
	}

	var _strings []string
	err = json.Unmarshal([]byte(_resp), &_strings)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 2, "msg": err.Error()})
	}

	var _msisdnStates []bean.MsisdnStates
	for index, _string := range _strings {
		var _msisdnState bean.MsisdnStates
		err = json.Unmarshal([]byte(_string), &_msisdnState)
		if err != nil {
			_msisdnState.Msisdn = _formMsisdns.Msisdns[index]
			_msisdnStates = append(_msisdnStates, _msisdnState)
			continue
		}
		_msisdnStates = append(_msisdnStates, _msisdnState)
	}

	c.JSON(http.StatusOK, _msisdnStates)
}

// GetMsisdnHistory GetMsisdnHistory
func GetMsisdnHistory(c *gin.Context) {

	_formMsisdn := bean.FormMsisdn{}
	c.BindJSON(&_formMsisdn)

	_encode := url.QueryEscape(_formMsisdn.Msisdn)

	_resp, err := blockchain.GetBCAPI().GetMsisdnHistory(_encode)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "msg": err.Error()})
	}

	if len(_resp) == 0 {
		c.JSON(200, gin.H{"errcode": 1, "msg": "response is empty"})
		return
	}

	var _histories []bean.TxHistory
	json.Unmarshal([]byte(_resp), &_histories)

	var _historyOuts []bean.TxHistoryOut
	for _, _history := range _histories {
		_string := strings.Replace(_history.Value, "\\", "", -1)

		var _historyOut bean.TxHistoryOut
		_historyOut.TxID = _history.TxID
		_historyOut.Timestamp = _history.Timestamp
		_historyOut.IsDelete = _history.IsDelete

		json.Unmarshal([]byte(_string), &_historyOut.Data)

		_historyOuts = append(_historyOuts, _historyOut)
	}

	c.JSON(http.StatusOK, _historyOuts)
}

// GetAssetInfo GetAssetInfo
func GetAssetInfo(c *gin.Context) {

	_formAsset := bean.FormAsset{}
	c.BindJSON(&_formAsset)

	_encode := url.QueryEscape(_formAsset.AssetID)

	fmt.Println(_encode)
	_resp, err := blockchain.GetBCAPI().GetAssetInfo(_encode)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "msg": err.Error()})
	}

	if len(_resp) == 0 {
		c.JSON(200, gin.H{"errcode": 1, "msg": "response is empty"})
	} else {
		c.String(http.StatusOK, _resp)
	}
}
