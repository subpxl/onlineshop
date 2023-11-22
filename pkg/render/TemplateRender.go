package render

import (
	"html/template"
	"net/http"
	"onlineshop/pkg/models"
)

type TemplateData struct {
	IsLoggedIn bool
}

type TemplateRenderHandler struct {
	Template     *template.Template
	TemplateData *TemplateData
}

func NewTemplateRenderHandler(TemplateDir string) *TemplateRenderHandler {
	templates := template.Must(template.ParseGlob(TemplateDir + "/*.html"))

	return &TemplateRenderHandler{
		Template: templates,
	}

}

func (tr *TemplateRenderHandler) Render(w http.ResponseWriter, r *http.Request, templateName string, data interface{}) {
	w.Header().Set("Content-Type", "text/html")

	dataWithCategories := struct {
		Categories []models.Category
		Data       interface{}
	}{
		Categories: models.GetAllCategories(),
		Data:       data,
	}

	if err := tr.Template.ExecuteTemplate(w, templateName, dataWithCategories); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
