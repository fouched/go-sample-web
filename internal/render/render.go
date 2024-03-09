package render

import (
	"fmt"
	"github.com/fouched/go-sample-web/internal/config"
	"github.com/fouched/go-sample-web/internal/models"
	"html/template"
	"net/http"
)

var pathToTemplates = "./templates"
var app *config.AppConfig

func NewRenderer(a *config.AppConfig) {
	app = a
}

func Template(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {

	td = AddDefaultData(td, r)

	parsedTemplate, _ := template.ParseFiles(pathToTemplates+tmpl, pathToTemplates+"/base.layout.tmpl")
	err := parsedTemplate.Execute(w, td)
	if err != nil {
		fmt.Println("Error parsing template", err)
		return
	}

}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {

	// add all session keys to the template data StringMap for illustration purposes
	// normally we would use session data much more precisely
	keys := app.Session.Keys(r.Context())
	for i := range keys {
		td.StringMap[keys[i]] = app.Session.GetString(r.Context(), keys[i])
	}

	return td
}
