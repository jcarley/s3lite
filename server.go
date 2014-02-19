package main

import (
  "net/http"
  "github.com/jcarley/s3lite/domain"
  "github.com/jcarley/s3lite/infrastructure"
  "github.com/jcarley/s3lite/controllers"
  "github.com/jcarley/s3lite/webservice"
  "github.com/codegangsta/martini"
)

func DB() martini.Handler {
  return func(c martini.Context) {
    var db domain.Database
    db = infrastructure.NewInMemoryDatabase()
    c.MapTo(db, (*domain.Database)(nil))
    c.Next()
  }
}

func BS() martini.Handler {
  return func(c martini.Context) {
    var db domain.BlobStorage
    bs = infrastructure.NewInMemoryBlobStorage()
    c.MapTo(bs, (*domain.BlobStorage)(nil))
    c.Next()
  }
}

func main() {
  m := martini.Classic()
  m.Use(DB())
  m.Use(BS())

  uploadController := controllers.NewUploadController()
  webservice.RegisterWebService(uploadController, m)

  http.ListenAndServe(":8080", m)
}

