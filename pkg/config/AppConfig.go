package config

import (
	"html/template"
	"onlineshop/pkg/render"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	InProduction bool
	Render       *render.TemplateRenderHandler
	Session      *scs.SessionManager
	TemplateData *template.Template
	RazorpayId   string
}

func NewAppCOnfig() *AppConfig {
	return &AppConfig{}
}
