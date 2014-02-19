package webservice

import (
  "fmt"
  "net/http"
  "github.com/jcarley/s3lite/domain"
  "github.com/codegangsta/martini"
)

// valid values for x-amz-acl:
  // private | public-read | public-read-write | authenticated-read | bucket-owner-read | bucket-owner-full-control

type WebService interface {
  InitiateMultipartUpload(req *http.Request, db domain.Database, blobStorage domain.BlobStorage) (int, string)
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

    // users GET    /users(.:format)                        users#index
          // POST   /users(.:format)                        users#create
 // new_user GET    /users/new(.:format)                    users#new
// edit_user GET    /users/:id/edit(.:format)               users#edit
     // user GET    /users/:id(.:format)                    users#show
          // PUT    /users/:id(.:format)                    users#update
          // DELETE /users/:id(.:format)                    users#destroy

