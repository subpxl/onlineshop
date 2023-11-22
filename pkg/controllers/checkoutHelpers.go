package controllers

import (
	"fmt"
	"net/http"
	"onlineshop/pkg/models"
	"strconv"
	"strings"
)

func createOrderItemsHelper(w http.ResponseWriter, r *http.Request, order models.Order, sh *ShopHandler) (int, int, int, error) {

	total := 0
	subtotal := 0
	tax := 0

	order_items := []models.OrderItem{}

	for key, values := range r.PostForm {

		for _, value := range values {

			if strings.HasPrefix(key, "quantity_") {

				indexStr := key[9:]

				productKey := "product_id_" + indexStr

				productValue := r.PostForm.Get(productKey)

				quantity, err := strconv.Atoi(value)

				if err != nil {

					fmt.Printf("Failed to parse quantity for %s: %v\n", key, err)

				} else {

					fmt.Printf("Quantity for %s: %d\n", productKey, quantity)
				}

				productID, err := strconv.Atoi(productValue)

				if err != nil {

					fmt.Printf("Failed to parse product ID for %s: %v\n", productKey, err)
					return 0, 0, 0, err
				} else {

					fmt.Printf("Product ID for %s: %d\n", productKey, productID)

					product := models.GetProductById(productID)

					subtotal += product.Price

					tax += product.Tax

					total = subtotal + tax
					orderItem := models.OrderItem{
						ProductID: productID,
						Quantity:  quantity,
						OrderID:   order.ID,
					}
					// models.CreateOrderItem(orderItem)
					//
					order_items = append(order_items, orderItem)

					// Save the changes to the database
					// models.
				}
			}
		}

	}
	for _, orderItem := range order_items {

		models.CreateOrderItem(orderItem)

	}
	return subtotal, tax, total, nil
}

func createOrderHelper(w http.ResponseWriter, r *http.Request, address models.Address, sh *ShopHandler) (models.Order, error) {
	total := 0
	subtotal := 0
	tax := 0
	var order models.Order
	order.Address = address
	order.Status = string(models.Pending)

	username := sh.App.Session.GetString(r.Context(), "user")
	if username == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	fmt.Println("username is ", username)
	user := models.GetUserByUsername(username)
	if user.ID == 0 {
		http.Error(w, "user nor dound", http.StatusInternalServerError)
	} else {
		order.User = user

	}
	order.User = user
	total = subtotal + tax
	order.SubTotal = subtotal
	order.Tax = tax
	order.Total = total

	orderResponse, err := models.CreateOrder(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return orderResponse, nil
}

func saveAddressHelper(r *http.Request, w http.ResponseWriter) (models.Address, error) {
	var address models.Address
	address_line_1 := r.FormValue("address_line_1")
	address_line_2 := r.FormValue("address_line_2")
	landmark := r.FormValue("landmark")
	country := r.FormValue("country")
	pincode := r.FormValue("pincode")

	first_name := r.FormValue("first_name")
	last_name := r.FormValue("last_name")
	phone := r.FormValue("phone")
	email := r.FormValue("email")

	address.Addressline1 = address_line_1
	address.Addressline2 = address_line_2
	address.Landmark = landmark
	address.Country = country
	address.Pincode = pincode

	address.FirstName = first_name
	address.LastName = last_name
	address.Phone = phone
	address.Email = email

	addressResponsse, err := models.CreateAddress(address)
	if err != nil {

		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return models.Address{}, err
	}
	return addressResponsse, nil
}
