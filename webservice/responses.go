package webservice

// <?xml version="1.0" encoding="UTF-8"?>
// <InitiateMultipartUploadResult xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
  // <Bucket>example-bucket</Bucket>
  // <Key>example-object</Key>
  // <UploadId>VXBsb2FkIElEIGZvciA2aWWpbmcncyBteS1tb3ZpZS5tMnRzIHVwbG9hZA</UploadId>
// </InitiateMultipartUploadResult>
type InitiateMultipartUploadResult struct {
  Bucket string
  Key string
  UploadId string
}
