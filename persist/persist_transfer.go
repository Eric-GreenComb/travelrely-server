package persist

import (
	"github.com/Eric-GreenComb/travelrely-server/bean"
)

// CreateTransfer CreateTransfer Persist
func (persist *Persist) CreateTransfer(transfer bean.Transfers) error {
	return persist.db.Create(&transfer).Error
}

// QueryTransfer QueryTransfer
func (persist *Persist) QueryTransfer(_search bean.FormSearchTx) ([]bean.Transfers, error) {
	var txs []bean.Transfers

	err := persist.db.Table("transfers").Where("tm >= ?", _search.Tm).Limit(_search.MaxNo).Offset(_search.Start).Find(&txs).Error

	return txs, err
}
