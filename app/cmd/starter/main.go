package main

import (
	"fmt"
	"github.com/ai-musics/go-starter/internal/greet"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"html/template"
	"time"
)

func main() {
	// Echo instance
	e := echo.New()

	// Create renderer
	gv := echoview.New(goview.Config{
		Root:      "views",
		Extension: ".html",
		Master:    "layouts/application",
		Funcs: template.FuncMap{
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
	})

	// Renderer
	e.Renderer = gv

	e.GET("/", hello)

	e.Logger.Fatal(e.Start(":3001"))
}

func hello(c echo.Context) error {
	fmt.Println("Go modules!")

	return c.Render(http.StatusOK, "index", echo.Map{
		"Title": greet.TheWorld(),
	})
}