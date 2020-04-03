package tempengine

import (
	"errors"
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo"
)

type TemplateRegistry struct {
	Templates map[string]*template.Template
}

// Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.Templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

func GetPages() []string {

	Pages := []string{
		"index.html",
		"jsonload.html",
		"csvload.html",
	}

	return Pages
}

func GetTemplates() map[string]*template.Template {

	templates := make(map[string]*template.Template)
	baseTemplate := "public/views/base.html"

	for _, page := range GetPages() {
		templates[page] = template.Must(template.ParseFiles("public/views/" + page, baseTemplate))
	}

	/*templates["index.html"] = template.Must(template.ParseFiles("public/views/index.html", baseTemplate))
	templates["jsonload.html"] = template.Must(template.ParseFiles("public/views/jsonload.html", baseTemplate))
	templates["csvload.html"] = template.Must(template.ParseFiles("public/views/csvload.html", baseTemplate))*/

	fmt.Println(templates)
	
	return templates
}
