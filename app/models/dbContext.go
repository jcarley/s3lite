package models

import (
	"database/sql"
	"log"

	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
)

type Database interface {
	GetUploadByUploadId(uploadId string) *Upload
	CreateUpload(upload *Upload) (uint, error)
}

type DbContext struct {
	Dbmap *gorp.DbMap
}

func NewDbContext() DbContext {
	db, err := sql.Open("mysql", "")
	if err != nil {
		log.Fatal(err)
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	return DbContext{Dbmap: dbmap}
}
