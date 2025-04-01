package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/klimenkokayot/avito-go/front/config"
	"github.com/klimenkokayot/avito-go/libs/logger"
)

type ViewHandler struct {
	templates *template.Template
	logger    logger.Logger
	cfg       *config.Config
}

func NewViewHandler(cfg *config.Config, logger logger.Logger) (*ViewHandler, error) {
	templateDir := filepath.Join("web", "template", "*.html")
	tmpl, err := template.ParseGlob(templateDir)
	if err != nil {
		return nil, ErrTemplateParse
	}
	return &ViewHandler{
		templates: tmpl,
		logger:    logger,
		cfg:       cfg,
	}, nil
}

func (a *ViewHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	err := a.templates.ExecuteTemplate(w, "auth_login.html", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a *ViewHandler) RegisterPage(w http.ResponseWriter, r *http.Request) {
	err := a.templates.ExecuteTemplate(w, "auth_register.html", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
