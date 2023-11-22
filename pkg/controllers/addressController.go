package controllers

import (
	"encoding/json"
	"net/http"
	"onlineshop/pkg/models"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/justinas/nosurf"
)

func (uh *UserHanlder) CreateAddress(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		var address models.Address

		first_name := r.FormValue("first_name")
		last_name := r.FormValue("last_name")
		phone := r.FormValue("phone")
		email := r.FormValue("email")

		address_line_1 := r.FormValue("address_line_1")
		address_line_2 := r.FormValue("address_line_2")
		landmark := r.FormValue("landmark")
		country := r.FormValue("country")
		pincode := r.FormValue("pincode")

		address.Addressline1 = address_line_1
		address.Addressline2 = address_line_2
		address.Landmark = landmark
		address.Country = country
		address.Pincode = pincode

		address.FirstName = first_name
		address.LastName = last_name
		address.Phone = phone
		address.Email = email

		resp, _ := models.CreateAddress(address)
		json.NewEncoder(w).Encode(resp)

	} else {

		token := nosurf.Token(r)
		data := struct {
			Token string
		}{Token: token}
		// sh.App.Render.Render(w, r, "CreateCategory.html", data)
		json.NewEncoder(w).Encode(data)

	}
}

func (uh *UserHanlder) UpdateAddress(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		var address models.Address

		first_name := r.FormValue("first_name")
		last_name := r.FormValue("last_name")
		phone := r.FormValue("phone")
		email := r.FormValue("email")
		address_line_1 := r.FormValue("address_line_1")
		address_line_2 := r.FormValue("address_line_2")
		landmark := r.FormValue("landmark")
		country := r.FormValue("country")
		pincode := r.FormValue("pincode")

		address.Addressline1 = address_line_1
		address.Addressline2 = address_line_2
		address.Landmark = landmark
		address.Country = country
		address.Pincode = pincode

		address.FirstName = first_name
		address.LastName = last_name
		address.Phone = phone
		address.Email = email

		resp := models.UpdateAddress(id, address)
		json.NewEncoder(w).Encode(resp)

	} else {

		token := nosurf.Token(r)
		data := struct {
			Token string
		}{Token: token}
		json.NewEncoder(w).Encode(data)

	}
}

func (uh *UserHanlder) DeleteAddress(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	resp := models.DeleteAddress(id)
	json.NewEncoder(w).Encode(resp)
}

func (uh *UserHanlder) GetAllAddresses(w http.ResponseWriter, r *http.Request) {
	addresses := models.GetAllAddresses()
	data := struct {
		Addresses []models.Address
	}{Addresses: addresses}
	json.NewEncoder(w).Encode(data)
}

func (uh *UserHanlder) GetAddress(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	data := models.GetAddressById(id)
	json.NewEncoder(w).Encode(data)
}
