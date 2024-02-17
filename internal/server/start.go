package server

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/nollidnosnhoj/simpbb/assets"
	"github.com/nollidnosnhoj/simpbb/internal/utils"
	"github.com/nollidnosnhoj/simpbb/internal/views"
)

func Start(context context.Context) {
	e := echo.New()

	fs := echo.MustSubFS(assets.Assets, "dist")
	e.StaticFS("/dist/", fs)

	e.GET("/", func (c echo.Context) error {
		view := views.IndexPage()
		return utils.RenderComponent(view, c)
	})

	server := http.Server{
		Addr: ":8080",
		Handler: e,
	}

	if err := server.ListenAndServe(); err != nil {
		e.Logger.Fatal(err)
	}

	<-context.Done()
	err := server.Close()
	if err != nil {
		e.Logger.Fatal(err)
	}
}