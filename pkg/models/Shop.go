package models

import (
	"errors"

	"gorm.io/gorm"
)

func RegisterShop(shop Shop) (bool, error) {
	resp := DB.Create(shop)
	if resp.Error != nil {
		return false, errors.New(resp.Error.Error())
	}
	return true, nil
}

// // user can add to cart
func AddToCart(cart Cart) Cart {
	DB.Create(&cart)
	return cart
}

// // user can place order
func PlaceOrder(order Order) *gorm.DB {
	resp := DB.Create(order)
	return resp
}
