package webservice

import (
  "fmt"
  "net/http"
  "github.com/jcarley/s3lite/domain"
  "github.com/codegangsta/martini"
)

type WebService interface {
  InitiateMultipartUpload(req *http.Request, db domain.Database) (int, string)
  UploadPart(params martini.Params, res http.ResponseWriter, req *http.Request) (int, string)
}

func RegisterWebService(webservice WebService, classicMartini *martini.ClassicMartini) {
  classicMartini.Get("/:id", func(params martini.Params, req *http.Request) string {
    fmt.Println(params)
    fmt.Println(req.Header.Get("Content-Disposition"))
    fmt.Println(req.URL.RawQuery)
    return "Hello"
  })
  classicMartini.Post("/**", webservice.InitiateMultipartUpload)
  classicMartini.Put("/**", webservice.UploadPart)
}
