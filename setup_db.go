package main

import (
  "database/sql"

  "github.com/jcarley/s3lite/domain"

  "github.com/coopernurse/gorp"
  _ "github.com/go-sql-driver/mysql"
)

func main() {
  db, err := sql.Open("mysql", "root@tcp/s3lite_development")
  if err != nil {
    panic(err)
  }

  dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
  t1 := dbmap.AddTableWithName(domain.Bucket{}, "buckets").SetKeys(true, "Id")
  t1.ColMap("Name").SetMaxSize(128)

  dbmap.DropTables()
  dbmap.CreateTablesIfNotExists()
}
