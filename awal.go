package main

import (
	"net/http"
    "html/template"
    "io"
    "ordertujuan/klass"
	
	"github.com/labstack/echo"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func Hello(c echo.Context) error {
    return c.Render(http.StatusOK, "hello", "World")
}

func main() {
	e := echo.New()
    e.Static("/", "staticfiles")
    
    
    renderer := &TemplateRenderer{
      templates: template.Must(template.ParseGlob("html/*.html")),
  }
  
    e.Renderer = renderer
    
    e.GET("/", klass.Awal)
    
    
	e.Logger.Fatal(e.Start(":8822"))
} 
