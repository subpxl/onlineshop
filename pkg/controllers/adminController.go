package controllers

import (
	"net/http"
	"onlineshop/pkg/config"
)

type AdminHandler struct {
	App *config.AppConfig
}

func NewAdminHandler(app *config.AppConfig) *AdminHandler {
	return &AdminHandler{
		App: app,
	}
}

// admin can login

// admin can register

// admin can registershop
// admin can crud product category
// admin can update order status
// admin can add shop details

func (ah *AdminHandler) Login(w http.ResponseWriter, r *http.Request)        {}
func (ah *AdminHandler) Register(w http.ResponseWriter, r *http.Request)     {}
func (ah *AdminHandler) RegisterShop(w http.ResponseWriter, r *http.Request) {}
func (ah *AdminHandler) UpdateShop(w http.ResponseWriter, r *http.Request)   {}
