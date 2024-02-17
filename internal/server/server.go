package server

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/nollidnosnhoj/simpbb/assets"
	"github.com/nollidnosnhoj/simpbb/internal/database"
	"github.com/nollidnosnhoj/simpbb/internal/utils"
	"github.com/nollidnosnhoj/simpbb/internal/views"
	"github.com/uptrace/bun"
)

type Server struct {
	Log *log.Logger
	Echo *echo.Echo
	Db *bun.DB
}

func NewServer() *Server {
	e := echo.New()

	logger := log.Default()

	fs := echo.MustSubFS(assets.Assets, "dist")
	e.StaticFS("/dist/", fs)

	// TODO: move this to controllers
	e.GET("/", func (c echo.Context) error {
		view := views.IndexPage()
		return utils.RenderComponent(view, c)
	})

	db := database.NewDb()

	return &Server{
		Log: logger,
		Echo: e,
		Db: db,
	}
}