package utilities

import (
  "net/http"
  "regexp"
  "strings"
)

var rxFilename = regexp.MustCompile(`filename=(.*)$`)
var rxBucket = regexp.MustCompile(`^([\S\w\-]+)\.s3`)
var rxKey = regexp.MustCompile(`^.*\/([a-zA-z\/]*)$`)

func ParseFilename(req *http.Request) string {
  matches := rxFilename.FindStringSubmatch(req.Header.Get("Content-Disposition"))

  filename := ""
  if len(matches) > 1 {
    filename = matches[1]
  }

  return filename
}

func ParseBucket(req *http.Request) string {
  matches := rxBucket.FindStringSubmatch(req.Host)

  bucket := "default"
  if len(matches) > 1 {
    bucket = matches[1]
  }
  return bucket
}

func ParseKey(req *http.Request) string {
  key := req.URL.Path

  if strings.Index(key, "/") == 0 {
    key = key[1:]
  }

  return key
}
