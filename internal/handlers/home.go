package handlers

import (
	"github.com/fouched/go-sample-web/internal/models"
	"github.com/fouched/go-sample-web/internal/render"
	"net/http"
)

func (m *HandlerConfig) Home(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	data["firstName"] = "Fouche"
	data["lastName"] = "du Preez"

	m.App.Session.Put(r.Context(), "SomeKey", "Hello from the session!")

	render.Template(w, r, "/home.page.tmpl", &models.TemplateData{
		StringMap: data,
	})
}
