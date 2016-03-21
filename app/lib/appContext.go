package lib

import (
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/jcarley/s3lite/app/models"
)

type AppContext struct {
	DbContext      models.DbContext
	Request        *http.Request
	Renderer       render.Render
	MartiniContext martini.Context
}
