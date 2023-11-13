package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
	"udemyCourse1/internal/config"
	"udemyCourse1/internal/models"

	"github.com/justinas/nosurf"
)

var app *config.AppConfig

// NewTemplates creates a new template cache
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	// get template cache from app config
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCashe()
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from cache")
	}

	buf := new(bytes.Buffer)

	// add default data to template data
	td = AddDefaultData(td, r)

	err := t.Execute(buf, td)
	if err != nil {
		log.Fatal(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}
}

// create a template cache
func CreateTemplateCashe() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all files *.page.tmpl from templates ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	// range over pages
	for _, page := range pages {
		// get file name
		name := filepath.Base(page)
		// parse page
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		// get base layout
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		// add to cache
		myCache[name] = ts
	}
	return myCache, nil
}