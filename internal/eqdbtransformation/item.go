package eqdbtransformation

import "dasheq/internal/eqdbobject"

type ItemMerchantList struct {
	MerchantID uint32
	ItemID     uint32
}

func JoinItemToMerchant(i *[]eqdbobject.Item) *[]ItemMerchantList {
	var iToM []ItemMerchantList
	return &iToM
}
