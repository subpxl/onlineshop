package controllers

import (
	"net/http"
	"onlineshop/pkg/models"
)

func (sh *ShopHandler) AddToCart(w http.ResponseWriter, r *http.Request) {

	var cart models.Cart

	data := struct {
		Cart models.Cart
	}{Cart: cart}

	sh.App.Render.Render(w, r, "cart.html", data)
}

// // user can add to cart
func (sh *ShopHandler) DeleteCart(w http.ResponseWriter, r *http.Request) {

	// product_name := chi.URLParam(r, "product_name")
	// product := models.GetProductByName(product_name)
	var cart models.Cart
	// cart.Products = append(cart.Products, product)
	resp := models.AddToCart(cart)
	sh.App.Render.Render(w, r, "cart.html", resp)
}

// // user can updaate cate
func (sh *ShopHandler) UpdateCart(w http.ResponseWriter, r *http.Request) {

	sh.App.Render.Render(w, r, "cart.html", nil)
}

// // user can updaate cate
func (sh *ShopHandler) CartPage(w http.ResponseWriter, r *http.Request) {
	sh.App.Render.Render(w, r, "cart.html", nil)
}
