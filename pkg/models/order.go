package models

func GetAllOrders() ([]Order, error) {
	var orders []Order

	if err := DB.
		Preload("User").
		Preload("OrderItems.Product").
		Preload("Address").
		Preload("Products").
		Find(&orders).
		Error; err != nil {
		return orders, err
	}

	return orders, nil
}

// // user can see Order page
func GetOrderById(id int) Order {
	var order Order
	DB.Preload("User").
		Preload("OrderItems.Product").
		Preload("Address").
		Preload("Products").Where("id = ?", id).First(&order)
	return order
}

func CreateOrder(order Order) (Order, error) {
	resp := DB.Create(&order)
	if resp.Error != nil {
		return Order{}, resp.Error
	}
	return order, nil
}
func UpdateOrder(id int, order Order) Order {
	DB.Where("id = ?", id).Updates(&order)
	return order
}
func DeleteOrder(id int) Order {
	var order Order
	DB.Where("id = ?", id).First(&order)
	DB.Delete(order)
	return order
}

func CreateOrderItem(order_item OrderItem) OrderItem {
	DB.Create(&order_item)
	return order_item
}
