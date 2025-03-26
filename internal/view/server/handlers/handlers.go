package view

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type ViewHandler struct {
	templates *template.Template
}

func NewAuthHandler() (*ViewHandler, error) {
	templateDir := filepath.Join("internal", "view", "web", "template", "*.html")
	tmpl, err := template.ParseGlob(templateDir)
	if err != nil {
		return nil, ErrTemplateParse
	}
	return &ViewHandler{
		tmpl,
	}, nil
}

func (a *ViewHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	err := a.templates.ExecuteTemplate(w, "auth_login.html", nil)
	if err != nil {
		log.Printf("Ошибка при вызове ExecuteTemplate: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a *ViewHandler) RegisterPage(w http.ResponseWriter, r *http.Request) {
	err := a.templates.ExecuteTemplate(w, "auth_register.html", nil)
	if err != nil {
		log.Printf("Ошибка при вызове ExecuteTemplate: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
