package main

import (
	"net/http"
	"onlineshop/pkg/config"
	"onlineshop/pkg/controllers"
	"onlineshop/pkg/middlewares"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func ShopRouter(app *config.AppConfig) *chi.Mux {
	// wiring
	r := chi.NewRouter()
	r.Use(app.Session.LoadAndSave)
	r.Use(middleware.Recoverer)
	r.Use(middlewares.NoSurfHandler)

	auth := middlewares.NewAuthMiddleware(app)
	uh := controllers.NewUserHandler(app)
	sh := controllers.NewShopHandler(app)

	fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	r.Handle("/static/*", fileServer)

	Uploadedimages := http.FileServer(http.Dir("./uploads/"))
	r.Handle("/uploads/*", http.StripPrefix("/uploads", Uploadedimages))

	r.Get("/", sh.HomePage)
	r.Get("/cart", sh.CartPage)
	r.Get("/registershop", sh.RegisterShop)
	r.Get("/store", sh.StorePage)
	r.Get("/search", sh.SearchPage)

	// customer route
	r.Get("/register", uh.Register)
	r.Post("/register", uh.Register)
	r.Get("/login", uh.Login)
	r.Post("/login", uh.Login)
	r.Get("/logout", uh.Logout)

	r.Get("/category/{category_name}", sh.StorePage)

	r.Get("/cart/{product_name}/{quantity}", sh.AddToCart)
	r.Get("/cart/update/{id}", sh.UpdateCart)
	r.Get("/cart/delete/{id}", sh.DeleteCart)

	r.Get("/product/{name}", sh.GetProduct)

	r.Get("/callback", sh.Callback)
	r.Post("/callback", sh.Callback)
	r.With(auth.AuthMiddleware).Get("/checkout", sh.Checkout)
	r.With(auth.AuthMiddleware).Post("/checkout", sh.Checkout)
	r.With(auth.AuthMiddleware).Get("/placeorder", sh.PlaceOrder)

	r.Get("/send-verification-link", uh.SendVerificationEmailHandler)
	r.Get("/verify/{token}", uh.VerifyEmail)

	AdminRoutes(r, sh)

	return r
}

func AdminRoutes(r *chi.Mux, sh *controllers.ShopHandler) {
	// secure this
	r.Get("/dashboard", sh.CustomerDashBoard)
	r.Get("/admin", sh.AdminDashBoard)

	r.Get("/admin/product", sh.GetAllProducts)
	r.Get("/admin/product/create", sh.CreateProduct)
	r.Post("/admin/product/create", sh.CreateProduct)
	r.Get("/admin/product/{id}", sh.GetProduct)
	r.Get("/admin/product/update/{id}", sh.UpdateProduct)
	r.Post("/admin/product/update/{id}", sh.UpdateProduct)
	r.Get("/admin/product/delete/{id}", sh.DeleteProduct)

	r.Get("/admin/category/create", sh.CreateCategory)
	r.Post("/admin/category/create", sh.CreateCategory)
	r.Get("/admin/category", sh.GetAllCategories)
	r.Get("/admin/category/{id}", sh.GetCategory)
	r.Get("/admin/category/update/{id}", sh.UpdateCategory)
	r.Post("/admin/category/update/{id}", sh.UpdateCategory)
	r.Get("/admin/category/delete/{id}", sh.DeleteCategory)

	r.Get("/admin/order", sh.GetAllOrders)
	r.Get("/admin/order/{id}", sh.GetOrderByID)
	r.Get("/admin/order/update/{id}", sh.UpdateOrder)
	r.Post("/admin/order/update/{id}", sh.UpdateOrder)
	r.Get("/admin/order/delete/{id}", sh.DeleteOrder)
	r.Get("/order/{id}", sh.GetOrderByID)

}
