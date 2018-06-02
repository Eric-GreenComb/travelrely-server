package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/travelrely-server/bean"
	"github.com/Eric-GreenComb/travelrely-server/persist"
)

// CreateTransfer CreateTransfer
func CreateTransfer(c *gin.Context) {
	var _formTransfers bean.FormTransfer
	c.BindJSON(&_formTransfers)

	var _tx bean.Transfers

	_tx.Tm = time.Now().Unix()
	_tx.Msisdn = _formTransfers.Msisdn
	_tx.BcuserID = _formTransfers.BCUserID
	_tx.AssetID = _formTransfers.AssetID
	_tx.Op = _formTransfers.Op

	err := persist.GetPersist().CreateTransfer(_tx)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "errmsg": "can't create tx."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": 0})
}

// QueryTransfer QueryTransfer
func QueryTransfer(c *gin.Context) {
	var _search bean.FormSearchTx
	c.BindJSON(&_search)

	_txs, err := persist.GetPersist().QueryTransfer(_search)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 1, "errmsg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": 0, "trans": _txs})
}
