package main

import (
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/jcarley/s3lite/app/lib"
	"github.com/jcarley/s3lite/app/models"
)

func PopulateAppContext(martiniContext martini.Context, w http.ResponseWriter, request *http.Request, renderer render.Render) {
	dbContext := models.NewDbContext()
	appContext := lib.AppContext{DbContext: dbContext, Request: request, Renderer: renderer, MartiniContext: martiniContext}

	martiniContext.Map(appContext)
}

func CloseDatabase(martiniContext martini.Context, appContext *lib.AppContext) {
	martiniContext.Next()
	appContext.DbContext.Dbmap.Db.Close()
}

func main() {
	m := martini.Classic()

	m.Use(PopulateAppContext)
	m.Use(CloseDatabase)

	http.ListenAndServe(":8080", m)
}
