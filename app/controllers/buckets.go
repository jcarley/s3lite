package buckets

import (
  "net/http"
  // "regexp"

  "github.com/codegangsta/martini"
  "github.com/jcarley/s3lite/domain"
  "github.com/jcarley/s3lite/utilities"
  // "github.com/jcarley/s3lite/encoding"
)

func RegisterWebService(classicMartini *martini.ClassicMartini) {
  classicMartini.Put("/", CreateBucket)
}

func CreateBucket(req *http.Request, db domain.Database) (int, string) {
  return 201, utilities.ParseBucket(req)
}
