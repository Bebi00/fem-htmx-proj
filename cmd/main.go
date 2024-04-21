package main

import (
	"html/template"
	"io"

	"github.com/Bebi00/htmx_intro/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

type Count struct {
	Count int
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	contacts := model.NewContactList(3)

	e.Renderer = newTemplate()
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", contacts)
	})
	e.POST("/contacts", func(c echo.Context) error {
		return c.Render(200, "count", contacts)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}
