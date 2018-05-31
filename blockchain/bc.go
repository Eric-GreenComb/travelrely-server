package blockchain

import ()

// BcAPI struct
type BcAPI struct {
}

// GBCAPI global GBCAPI
var GBCAPI *BcAPI

// GetBCAPI Get BcAPI
func GetBCAPI() *BcAPI {

	if GBCAPI != nil {
		return GBCAPI
	}

	GBCAPI = new(BcAPI)
	return GBCAPI
}
