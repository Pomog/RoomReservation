package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var tc = make(map[string]*template.Template)

func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	//check if template is already parsed
	_, ok := tc[t]
	if !ok {
		log.Println("creating template and adding to cache")
		err = createTemplateCashe(t)
		if err != nil {
			log.Println("error creating template cache", err)
		}
	} else {
		log.Println("template parsed")
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("error creating template cache", err)
	}
}

func createTemplateCashe(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	tc[t] = tmpl
	return nil
}
