package main

import (
	"fmt"
	"net/http"
	"onlineshop/pkg/config"
	"onlineshop/pkg/models"
	"onlineshop/pkg/render"
)

func init() {
	models.ConnectToDatabase()
	models.SyncDB()
}

func main() {

	session := getSession()
	renderer := render.NewTemplateRenderHandler("views")

	app := config.AppConfig{
		Session:    session,
		RazorpayId: "rzp_test_pDqm2g3OXuOBnj",
		Render:     renderer,
	}

	// start app
	fmt.Println("running sesrver on 8000")
	http.ListenAndServe(":8000", ShopRouter(&app))
}
