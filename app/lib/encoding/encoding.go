package encoding

import (
  "bytes"
  "encoding/xml"
)

// An Encoder implements an encoding format of values to be sent as response to
// requests on the API endpoints.
type Encoder interface {
  Encode(v ...interface{}) (string, error)
}

type XmlEncoder struct{}

// xmlEncoder is an Encoder that produces XML-formatted responses.
func (_ XmlEncoder) Encode(v ...interface{}) (string, error) {
  var buf bytes.Buffer
  if _, err := buf.Write([]byte(xml.Header)); err != nil {
    return "", err
  }
  b, err := xml.Marshal(v)
  if err != nil {
    return "", err
  }
  if _, err := buf.Write(b); err != nil {
    return "", err
  }
  return buf.String(), nil
}
