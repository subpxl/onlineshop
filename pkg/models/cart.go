package models

func GetAllCarts() []Cart {
	var carts []Cart
	DB.Find(&carts)
	return carts
}

// // user can see Cart page
func GetCartById(id int) Cart {
	var cart Cart
	DB.Where("id = ?", id).First(&cart)
	return cart
}

func GetCartByName(name string) Cart {
	var cart Cart
	DB.Where("name = ?", name).First(&cart)
	return cart
}

func CreateCart(cart Cart) Cart {
	DB.Create(&cart)
	return cart
}
func UpdateCart(id int, cart Cart) Cart {
	DB.Where("id=?", id).Updates(&cart)
	return cart
}
func DeleteCart(id int) Cart {
	var Cart Cart
	DB.Where("id= ?", id).First(&Cart)
	DB.Delete(Cart)
	return Cart
}

func CreateCartItem(cartItem CartItem) CartItem {
	DB.Create(&cartItem)
	return cartItem
}
