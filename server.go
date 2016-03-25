package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jcarley/s3lite/controllers"

	r "github.com/dancannon/gorethink"
)

func main() {

	//TODO:  Introduce an AppContext

	session := connectToDB()

	router := mux.NewRouter()

	uploadController := controllers.NewUploadController(session)
	uploadController.Register(router)

	bucketController := controllers.NewBucketController(session)
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
