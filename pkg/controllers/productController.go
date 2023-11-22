package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"onlineshop/pkg/models"
	"os"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/justinas/nosurf"
)

// // user can see products list
func (sh *ShopHandler) ProductList(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts()
	data := struct {
		Products []models.Product
	}{
		Products: products,
	}
	sh.App.Render.Render(w, r, "productList.html", data)
}

func (sh *ShopHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Println("i am called prdocut create")
		// Parse the form data
		err := r.ParseMultipartForm(10 << 20) // 10MB limit
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		// Extract form values
		name := r.FormValue("name")
		price, err := strconv.Atoi(r.FormValue("price"))
		if err != nil {
			http.Error(w, "Invalid price value", http.StatusBadRequest)
			return
		}
		salePrice, err := strconv.Atoi(r.FormValue("sale_price"))
		if err != nil {
			http.Error(w, "Invalid sale price value", http.StatusBadRequest)
			return
		}
		categoryID, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			http.Error(w, "Invalid category ID value", http.StatusBadRequest)
			return
		}
		description := r.FormValue("description")

		// Handle file upload
		file, handler, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "Error retrieving the uploaded file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		uploadDir := "./uploads/"
		// Create the directory if it doesn't exist
		if err := os.MkdirAll(uploadDir, 0755); err != nil {
			http.Error(w, "Error creating upload directory", http.StatusInternalServerError)
			return
		}

		// Create a new file in the upload directory
		// fileLocation := filepath.Join(uploadDir, handler.Filename)
		fileLocation := fmt.Sprintf("%s/%s", uploadDir, handler.Filename)
		newFile, err := os.Create(fileLocation)
		if err != nil {
			http.Error(w, "Error creating the file", http.StatusInternalServerError)
			return
		}
		defer newFile.Close()

		// Copy the uploaded file data to the new file
		_, err = io.Copy(newFile, file)
		if err != nil {
			http.Error(w, "Error copying file data", http.StatusInternalServerError)
			return
		}

		// Create a Product instance with the extracted data
		product := models.Product{
			Name:        name,
			Price:       price,
			SalePrice:   salePrice,
			CategoryID:  categoryID,
			Description: description,
			Image:       fileLocation, // Set the image file path
		}

		// Call the function to create the product
		models.CreateProduct(product)
		// json.NewEncoder(w).Encode(resp)
		http.Redirect(w, r, "/admin/product", http.StatusSeeOther)

	} else {
		// Handle the GET request, rendering a form
		token := nosurf.Token(r)
		data := struct {
			Token string
		}{Token: token}
		// fmt.Println("token iss", token)
		// json.NewEncoder(w).Encode(data)
		sh.App.Render.Render(w, r, "adminCreateProduct.html", data)
	}
}

func (sh *ShopHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		fmt.Println("i am called prdocut create")
		// Parse the form data
		err := r.ParseMultipartForm(10 << 20) // 10MB limit
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		// Extract form values
		name := r.FormValue("name")
		price, err := strconv.Atoi(r.FormValue("price"))
		if err != nil {
			http.Error(w, "Invalid price value", http.StatusBadRequest)
			return
		}
		salePrice, err := strconv.Atoi(r.FormValue("sale_price"))
		if err != nil {
			http.Error(w, "Invalid sale price value", http.StatusBadRequest)
			return
		}
		categoryID, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			http.Error(w, "Invalid category ID value", http.StatusBadRequest)
			return
		}
		description := r.FormValue("description")

		// Handle file upload
		file, handler, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "Error retrieving the uploaded file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		uploadDir := "./uploads/"
		// Create the directory if it doesn't exist
		if err := os.MkdirAll(uploadDir, 0755); err != nil {
			http.Error(w, "Error creating upload directory", http.StatusInternalServerError)
			return
		}

		// Create a new file in the upload directory
		// fileLocation := filepath.Join(uploadDir, handler.Filename)
		fileLocation := fmt.Sprintf("%s/%s", uploadDir, handler.Filename)
		newFile, err := os.Create(fileLocation)
		if err != nil {
			http.Error(w, "Error creating the file", http.StatusInternalServerError)
			return
		}
		defer newFile.Close()

		// Copy the uploaded file data to the new file
		_, err = io.Copy(newFile, file)
		if err != nil {
			http.Error(w, "Error copying file data", http.StatusInternalServerError)
			return
		}

		// Create a Product instance with the extracted data
		product := models.Product{
			Name:        name,
			Price:       price,
			SalePrice:   salePrice,
			CategoryID:  categoryID,
			Description: description,
			Image:       fileLocation, // Set the image file path
		}

		// Call the function to create the product
		models.UpdateProduct(id, product)
		// json.NewEncoder(w).Encode(resp)
		http.Redirect(w, r, "/admin/product", http.StatusSeeOther)

	} else {
		// Handle the GET request, rendering a form
		token := nosurf.Token(r)
		data := struct {
			Token string
		}{Token: token}
		// fmt.Println("token iss", token)
		// json.NewEncoder(w).Encode(data)
		sh.App.Render.Render(w, r, "adminCreateProduct.html", data)
	}
}

func (sh *ShopHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts()
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
	data := struct {
		Products []models.Product
	}{Products: products}
	// json.NewEncoder(w).Encode(data)
	sh.App.Render.Render(w, r, "adminProducts.html", data)

}

func (sh *ShopHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	link := chi.URLParam(r, "name")
	product := models.GetProductByName(link)
	data := struct {
		Product models.Product
	}{Product: product}
	sh.App.Render.Render(w, r, "productPage.html", data)
	// json.NewEncoder(w).Encode(data)
}

func (sh *ShopHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	resp := models.DeleteProduct(id)
	json.NewEncoder(w).Encode(resp)

}
