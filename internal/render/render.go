package render

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
	"udemyCourse1/internal/config"
	"udemyCourse1/internal/models"

	"github.com/justinas/nosurf"
)

var app *config.AppConfig
var pathToTemplates = "./templates"
var functions = template.FuncMap{}

// NewRenderer creates a new template cache
func NewRenderer(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.CSRFToken = nosurf.Token(r)
	return td
}

func Template(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) error {
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
		return errors.New("could not get template from cache")
	}

	buf := new(bytes.Buffer)

	// add default data to template data
	td = AddDefaultData(td, r)

	err := t.Execute(buf, td)
	if err != nil {
		return err
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		return err
	}

	return nil
}

// create a template cache
func CreateTemplateCashe() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all files *.page.tmpl from templates ./templates
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return myCache, err
	}
	// range over pages
	for _, page := range pages {
		// get file name
		name := filepath.Base(page)
		// parse page
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		// get base layout
		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}
		// add to cache
		myCache[name] = ts
	}
	return myCache, nil
}
