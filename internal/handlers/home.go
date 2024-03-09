package handlers

import (
	"github.com/fouched/go-sample-web/internal/models"
	"github.com/fouched/go-sample-web/internal/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	data["firstName"] = "Fouche"
	data["lastName"] = "du Preez"

	render.Template(w, r, "/home.page.tmpl", &models.TemplateData{
		StringMap: data,
	})
}
