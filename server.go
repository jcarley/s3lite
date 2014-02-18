package main

import (
  "net/http"
  "github.com/jcarley/s3lite/domain"
  "github.com/jcarley/s3lite/controllers"
  "github.com/jcarley/s3lite/webservice"
  "github.com/codegangsta/martini"
)

func DB() martini.Handler {
  return func(c martini.Context) {
    var db domain.Database
    db = domain.NewInMemoryDatabase()
    c.MapTo(db, (*domain.Database)(nil))
    c.Next()
  }
}

func main() {
  m := martini.Classic()
  m.Use(DB())

  uploadController := controllers.NewUploadController()
  webservice.RegisterWebService(uploadController, m)

  http.ListenAndServe(":8080", m)
}

