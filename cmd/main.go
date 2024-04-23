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

type IndexPageData struct {
	ContactList       model.ContactList
	CreateContactData model.CreateContactData
}

func newIndexPageData() IndexPageData {
	return IndexPageData{
		ContactList:       model.NewContactList(3),
		CreateContactData: model.NewCreateContactData(),
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	indexPage := newIndexPageData()

	e.Renderer = newTemplate()
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", indexPage)
	})
	e.POST("/contacts", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")

		if model.ExistsEmail(email, indexPage.ContactList.Contacts) {
			createContactData := model.NewCreateContactData()
			createContactData.Values["name"] = name
			createContactData.Values["email"] = email
			createContactData.Errors["email"] = "Email already exists"

			return c.Render(422, "create-contact", createContactData)
		}

		contact := model.NewContact(name, email)
		indexPage.ContactList.Contacts = append(indexPage.ContactList.Contacts, contact)

		c.Render(200, "create-contact", model.NewCreateContactData())
		return c.Render(200, "oob-contact", contact)
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
