package controllers

import (
	"encoding/json"
	"net/http"
	"onlineshop/pkg/models"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/justinas/nosurf"
)

// admin can crud product category
func (sh *ShopHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var category models.Category
		name := r.FormValue("name")
		category.Name = name
		models.CreateCategory(category)
		// json.NewEncoder(w).Encode(resp)
		// sh.App.Render.Render(w, r, "adminCategory.html", nil)
		http.Redirect(w, r, "/admin/category", http.StatusSeeOther)
	} else {

		token := nosurf.Token(r)
		data := struct {
			Token string
		}{Token: token}
		sh.App.Render.Render(w, r, "adminCreateCategory.html", data)

	}
}

func (sh *ShopHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		var category models.Category
		name := r.FormValue("name")
		id, _ := strconv.Atoi(r.FormValue("id"))
		category.Name = name
		models.UpdateCategory(id, category)
		// json.NewEncoder(w).Encode(resp)
		http.Redirect(w, r, "/admin/category", http.StatusSeeOther)

	} else {

		token := nosurf.Token(r)
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		category := models.GetCategoryById(id)
		data := struct {
			Category models.Category
			Token    string
		}{Token: token, Category: category}
		sh.App.Render.Render(w, r, "adminCreateCategory.html", data)

	}
}

func (sh *ShopHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {

	// id, _ := strconv.Atoi(r.FormValue("id"))
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	models.DeleteCategory(id)
	// json.NewEncoder(w).Encode(resp)
	http.Redirect(w, r, "/admin/category", http.StatusSeeOther)
}

func (sh *ShopHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories := models.GetAllCategories()
	data := struct {
		Categories []models.Category
	}{Categories: categories}
	sh.App.Render.Render(w, r, "adminCategories.html", data)

}

func (sh *ShopHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	data := models.GetCategoryById(id)
	json.NewEncoder(w).Encode(data)
}
