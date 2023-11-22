package controllers

import (
	"fmt"
	"net/http"
	"onlineshop/pkg/config"
	"onlineshop/pkg/models"

	"github.com/justinas/nosurf"
)

type UserHanlder struct {
	App *config.AppConfig
}

func NewUserHandler(app *config.AppConfig) *UserHanlder {
	return &UserHanlder{
		App: app,
	}
}

func (uh *UserHanlder) Register(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		fmt.Println("register page")
		username := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("email")
		city := r.FormValue("city")
		country := r.FormValue("country")
		firstName := r.FormValue("firstName")
		lastName := r.FormValue("lastName")
		gender := r.FormValue("gender")

		var user models.User
		user.City = city
		user.Country = country
		user.Email = email
		user.Gender = gender
		user.FirstName = firstName
		user.LastName = lastName
		user.Username = username
		user.Password = password
		user.Name = firstName + " " + lastName

		response, err := models.UserRegister(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		uh.App.Session.Put(r.Context(), "status", response)
		http.Redirect(w, r, "/login", http.StatusSeeOther)

	} else {
		token := nosurf.Token(r)
		data := struct {
			Token string
		}{Token: token}
		uh.App.Render.Render(w, r, "register.html", data)
	}
}

// user can login or login
func (uh *UserHanlder) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Println("login page")
		username := r.FormValue("username")
		password := r.FormValue("password")

		fmt.Println("username password", username, "  ", password)
		if password == "" || username == "" {
			http.Error(w, "empty password username not allowrd", http.StatusInternalServerError)
		}
		status, err := models.UserLogin(username, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// successfully logged in
		if status {
			uh.App.Session.Put(r.Context(), "user", username)

			role := models.GetUserByUsername(username).RoleID
			fmt.Println("user role is ", role)
			if role == 1 {
				http.Redirect(w, r, "/", http.StatusSeeOther)

			} else if role == 2 {
				http.Redirect(w, r, "/admin", http.StatusSeeOther)
			}

		}
	} else {
		token := nosurf.Token(r)
		data := struct {
			Token string
		}{Token: token}
		uh.App.Render.Render(w, r, "login.html", data)
	}

}
func (uh *UserHanlder) Logout(w http.ResponseWriter, r *http.Request) {
	uh.App.Session.Destroy(r.Context())
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func (uh *UserHanlder) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	// uh.App.Session.Destroy(r.Context())
	// http.Redirect(w, r, "/", http.StatusSeeOther)
	w.Write([]byte("to apply forgot password here"))
}
