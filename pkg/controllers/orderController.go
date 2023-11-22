package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"onlineshop/pkg/models"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/justinas/nosurf"
)

// admin can update order status

func (sh *ShopHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var order models.Order
		status := r.FormValue("status")
		// id, _ := strconv.Atoi(r.FormValue("id"))
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))

		order.Status = status

		models.UpdateOrder(id, order)
		// json.NewEncoder(w).Encode(resp)
		http.Redirect(w, r, "/admin/order", http.StatusSeeOther)

	} else {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))

		order := models.GetOrderById(id)
		token := nosurf.Token(r)
		data := struct {
			Token string
			Order models.Order
		}{Token: token, Order: order}
		sh.App.Render.Render(w, r, "adminCreateOrder.html", data)

	}
}

func (sh *ShopHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))

	resp := models.DeleteCart(id)
	json.NewEncoder(w).Encode(resp)
	http.Redirect(w, r, "/admin/order", http.StatusSeeOther)
}

func (sh *ShopHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	categories := models.GetAllCategories()
	data := struct {
		Categories []models.Category
	}{Categories: categories}
	sh.App.Render.Render(w, r, "category.html", data)

}

func (sh *ShopHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	fmt.Println("id is ", id)
	order := models.GetOrderById(id)

	data := struct {
		Order models.Order
	}{Order: order}
	sh.App.Render.Render(w, r, "orderPage.html", data)

	// json.NewEncoder(w).Encode(data)
}
