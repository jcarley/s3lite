package main

import (
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/jcarley/s3lite/buckets"
	"github.com/jcarley/s3lite/domain"
	"github.com/jcarley/s3lite/infrastructure"
)

func SetupDB() *domain.Database {
	// db, err := sql.Open("mysql", "user:password@/dbname")
	// PanicIf(err)
	// return db
	db := infrastructure.NewInMemoryDatabase()
	return &db.(*domain.Database)
}

func SetupBlobStorage() *domain.BlobStorage {
	bs := infrastructure.NewInMemoryBlobStorage()
	return &bs
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	m := martini.Classic()
	m.Map(SetupDB())
	m.Map(SetupBlobStorage())

	buckets.RegisterWebService(m)

	http.ListenAndServe(":8080", m)
}
