package main

import (
  "net/http"

  "github.com/codegangsta/martini"
  "github.com/jcarley/s3lite/buckets"
  "github.com/jcarley/s3lite/domain"
  "github.com/jcarley/s3lite/infrastructure"
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
    var bs domain.BlobStorage
    bs = infrastructure.NewInMemoryBlobStorage()
    c.MapTo(bs, (*domain.BlobStorage)(nil))
    c.Next()
  }
}

func main() {
  m := martini.Classic()
  m.Use(DB())
  m.Use(BS())

  buckets.RegisterWebService(m)

  http.ListenAndServe(":8080", m)
}
