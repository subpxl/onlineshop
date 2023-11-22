package models

// func ProductList(w http.ResponseWriter, r *http.Request)
// // user can see products list
func GetAllPayments() []Payment {
	var payments []Payment
	DB.Find(&payments)
	return payments
}

// // user can see product page
func GetPaymentById(id int) Payment {
	var payment Payment
	DB.Where("id = ?", id).First(&payment)
	return payment
}

func CreatePayment(payment Payment) (Payment, error) {

	respone := DB.Create(&payment)
	if respone.Error != nil {
		return Payment{}, respone.Error
	}
	return payment, nil
}

func UpdatePayment(id int, payment Payment) Payment {
	DB.Where("id = ?", id).Updates(&payment)
	return payment
}
