package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/travelrely-server/blockchain"
	"github.com/Eric-GreenComb/travelrely-server/config"
)

/////////////////////////////////////////////////////////////////
// BlockChain Area
/////////////////////////////////////////////////////////////////

// QueryChainHeight QueryChainHeight
func QueryChainHeight(c *gin.Context) {

	_resp, err := blockchain.GetBCAPI().QueryChainHeight(config.FabricConfig.APIChannel)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "msg": err.Error()})
	}

	if len(_resp) == 0 {
		c.JSON(200, gin.H{"errcode": 1, "msg": "response is empty"})
	} else {
		c.String(http.StatusOK, _resp)
	}
}
