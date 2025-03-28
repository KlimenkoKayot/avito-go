package front

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

type ViewHandler struct {
	templates *template.Template
}

func NewViewHandler() (*ViewHandler, error) {
	logrus.Info("Инициализация ViewHandler`a.")
	templateDir := filepath.Join("web", "template", "*.html")
	tmpl, err := template.ParseGlob(templateDir)
	if err != nil {
		logrus.Error("Ошибка при парсинге шаблонов в ViewHandler`e.")
		return nil, ErrTemplateParse
	}
	logrus.Debug("Успешно создан ViewHandler.")
	return &ViewHandler{
		tmpl,
	}, nil
}

func (a *ViewHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	logrus.Debug("Вызов страницы LoginPage.")
	err := a.templates.ExecuteTemplate(w, "auth_login.html", nil)
	if err != nil {
		logrus.Errorf("Ошибка при вызове ExecuteTemplate: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a *ViewHandler) RegisterPage(w http.ResponseWriter, r *http.Request) {
	logrus.Debug("Вызов страницы с RegisterPage.")
	err := a.templates.ExecuteTemplate(w, "auth_register.html", nil)
	if err != nil {
		logrus.Errorf("Ошибка при вызове ExecuteTemplate: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
