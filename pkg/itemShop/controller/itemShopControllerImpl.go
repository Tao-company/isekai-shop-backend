package controller

import (
	_itemShopservice "github.com/rratchapol/isekai-shop-api/pkg/itemShop/service"
)

type ItemShopControllerImpl struct {
	itemShopService _itemShopservice.ItemShopService
}

func NewItemShopControllerImpl(
	itemShopService _itemShopservice.ItemShopService,
) ItemShopController {
	return &ItemShopControllerImpl{itemShopService}
}
