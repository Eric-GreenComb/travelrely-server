package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/travelrely-server/bean"
	"github.com/Eric-GreenComb/travelrely-server/tools"
)

// RegisterUser RegisterUser
func RegisterUser(c *gin.Context) {

	_formUserRegister := bean.FormUserRegister{}
	c.BindJSON(&_formUserRegister)

	if _formUserRegister.CMD != "user_register" {
		c.JSON(http.StatusOK, gin.H{"result": 10, "errmsg": "非法访问"})
		return
	}
	// eth key
	_key, err := tools.Ks.NewKey()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 11, "msg": err.Error()})
		return
	}

	_pem, err := tools.GenPemPrivateKey(1024)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"result": 11, "errmsg": "gen pem error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": 0, "bcuser_id": _key.Address, "bcuser_key": _pem})
}

// UnRegisterUser UnRegisterUser
func UnRegisterUser(c *gin.Context) {

	_formUserUnRegister := bean.FormUserUnRegister{}
	c.BindJSON(&_formUserUnRegister)

	if _formUserUnRegister.CMD != "user_unregister" {
		c.JSON(http.StatusOK, gin.H{"result": 10, "errmsg": "非法访问"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": 0})
}
