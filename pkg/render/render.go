package render

import (
	"bytes"
	"github.com/code-chimp/gobookings/pkg/config"
	"github.com/code-chimp/gobookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// NewTemplates sets the config for the applications package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData add common data to templates
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders template found at t to ResponseWriter w using
// html/template
func RenderTemplate(w http.ResponseWriter, t string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache && len(app.TemplateCache) > 0 {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	tmpl, ok := tc[t]
	if !ok {
		log.Fatal("could not retrieve template from template cache")
	}

	// trap to verify template hydrated correctly
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)
	err := tmpl.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)

	// gather up all of the layouts
	layouts, err := filepath.Glob("./templates/*.layout.gohtml")
	if err != nil {
		return cache, err
	}

	// gather up all of the pages
	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		if len(layouts) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return cache, err
			}
		}

		cache[name] = ts
	}

	return cache, nil
}
