package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jcarley/s3lite/controllers"
	"github.com/jcarley/s3lite/domain"
	"github.com/jcarley/s3lite/services"

	r "github.com/dancannon/gorethink"
)

func main() {

	session := connectToDB()
	rethinkDatastore := domain.NewRethinkDatastore(session)

	router := mux.NewRouter()

	uploadService := services.NewUploadService(rethinkDatastore)
	bucketService := services.NewBucketService(rethinkDatastore)

	uploadController := controllers.NewUploadController(uploadService)
	uploadController.Register(router)

	bucketController := controllers.NewBucketController(bucketService)
	bucketController.Register(router)

	n := negroni.Classic()
	n.UseHandler(router)

	http.ListenAndServe(":8080", n)
}

func connectToDB() *r.Session {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "test",
		MaxIdle:  10,
		MaxOpen:  10,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	return session
}
