Start the server
---------------
go run server.go


Valid values for x-amz-acl
--------------------------
private
public-read
public-read-write
authenticated-read
bucket-owner-read
bucket-owner-full-control

Sample requests
---------------
-- Initiate a new multipart upload
curl -i -X "POST" --header "Content-Disposition: attachment; filename=foobar.mov" "http://localhost:8080/path/to/my/object?uploads"


