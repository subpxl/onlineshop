package controllers

import (
	"fmt"
	"net/http"
	"onlineshop/pkg/config"
	"onlineshop/pkg/models"
	"strconv"
)

type ShopHandler struct {
	App *config.AppConfig
}

func NewShopHandler(app *config.AppConfig) *ShopHandler {
	return &ShopHandler{
		App: app,
	}
}

func (sh *ShopHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts()
	categories := models.GetAllCategories()
	data := struct {
		Products   []models.Product
		Categories []models.Category
	}{Products: products, Categories: categories}
	sh.App.Render.Render(w, r, "index.html", data)
}

func (sh *ShopHandler) RegisterShop(w http.ResponseWriter, r *http.Request) {
	sh.App.Render.Render(w, r, "registerShop.html", nil)
}

func (sh *ShopHandler) StorePage(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts()
	data := struct {
		Products []models.Product
	}{
		Products: products,
	}
	sh.App.Render.Render(w, r, "store.html", data)
}

// // uesr can checkout
func (sh *ShopHandler) SearchPage(w http.ResponseWriter, r *http.Request) {

	sh.App.Render.Render(w, r, "search.html", nil)
}

// // uesr can checkout
func (sh *ShopHandler) Checkout(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}
		// saving address
		address, err := saveAddressHelper(r, w)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		order, err := createOrderHelper(w, r, address, sh)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// fmt.Println("order is ", order)

		sutbotal, tax, total, err := createOrderItemsHelper(w, r, order, sh)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(sutbotal, tax, total)
		// json.NewEncoder(w).Encode(orderItems)
		// var updatedUser User
		updatedorder := models.GetOrderById(order.ID)
		updatedorder.Total = total
		updatedorder.Tax = tax
		updatedorder.SubTotal = sutbotal
		fmt.Println("values are", updatedorder.Total, updatedorder.SubTotal, updatedorder.Tax)

		orderAfterTotal := models.UpdateOrder(order.ID, updatedorder)

		paymentResponse, err := ProcessPaymentHelper(w, r, orderAfterTotal)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// patment
		fmt.Println(paymentResponse)

		csrfToken := sh.App.Session.Token(r.Context())
		data := struct {
			Body       map[string]interface{}
			RazorpayId string
			User       models.User
			Order      models.Order
			Token      string
		}{Body: paymentResponse, RazorpayId: sh.App.RazorpayId, User: order.User, Order: updatedorder, Token: csrfToken}

		sh.App.Render.Render(w, r, "placeOrder.html", data)

	} else {
		sh.App.Render.Render(w, r, "checkout.html", nil)
	}

}

func (sh *ShopHandler) Callback(w http.ResponseWriter, r *http.Request) {
	// if r.Method == http.MethodPost {
	fmt.Println("Callback called........................")

	r.URL.Query().Get("version")

	signature := r.URL.Query().Get("razorpay_signature")
	payment_id := r.URL.Query().Get("razorpay_payment_id")
	razorpay_order_id := r.URL.Query().Get("razorpay_order_id")
	user_id, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
	order_id, _ := strconv.Atoi(r.URL.Query().Get("order_id"))

	fmt.Println("detais are sinature ", signature)
	fmt.Println("detais are paymetnid ", payment_id)
	fmt.Println("detais are user id  ", user_id)
	fmt.Println("detais are order id ", order_id)
	payment := models.Payment{OrderID: order_id, UserID: user_id, RazorpaySignature: signature, RazorpayPaymentID: payment_id, RazorpayOrderID: razorpay_order_id}
	paymentresponse, err := models.CreatePayment(payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	VerifyPayment := RazorPaymentVerification(signature, razorpay_order_id, payment_id)
	if VerifyPayment == nil {
		var neworder models.Order
		neworder.Status = string(models.Processing)
		var paymentStatus models.Payment
		paymentStatus.Status = models.Success
		models.UpdatePayment(paymentresponse.ID, paymentStatus)
		models.UpdateOrder(order_id, neworder)

	}

	myOrder := models.GetOrderById(order_id)
	data := struct {
		Payment models.Payment
		Order   models.Order
	}{Payment: paymentresponse, Order: myOrder}
	sh.App.Render.Render(w, r, "orderStatus.html", data)

	// json.NewEncoder(w).Encode(response)

}

// // user can place order
func (sh *ShopHandler) PlaceOrder(w http.ResponseWriter, r *http.Request) {

	sh.App.Render.Render(w, r, "placeOrder.html", nil)
}

// // user can see his dashboard for orders and status
func (sh *ShopHandler) CustomerDashBoard(w http.ResponseWriter, r *http.Request) {
	user := models.GetUserByUsername("lolo")
	orders, err := models.GetAllOrders()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	data := struct {
		User   models.User
		Orders []models.Order
	}{User: user, Orders: orders}
	// json.NewEncoder(w).Encode(data)
	sh.App.Render.Render(w, r, "customerDashboard.html", data)
}

// // user can see his dashboard for orders and status
func (sh *ShopHandler) AdminDashBoard(w http.ResponseWriter, r *http.Request) {

	orders, err := models.GetAllOrders()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	data := struct {
		Orders []models.Order
	}{Orders: orders}
	sh.App.Render.Render(w, r, "adminDashboard.html", data)
}

// admin can add shop details
func (sh *ShopHandler) CreateShop(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

	} else {
		sh.App.Render.Render(w, r, "registerShop.html", nil)

	}
}
func (sh *ShopHandler) UpdateShop(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

	} else {
		sh.App.Render.Render(w, r, "registerShop.html", nil)

	}
}
