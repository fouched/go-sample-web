package render

import (
	"fmt"
	"github.com/fouched/go-sample-web/internal/models"
	"html/template"
	"net/http"
)

var pathToTemplates = "./templates"

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
	return td
}
