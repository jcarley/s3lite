package infrastructure

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jcarley/s3lite/domain"
)

type MySQLDatabase struct {
	uploads map[uint]*domain.Upload
	seq     uint
}

func NewMySQLDatabase() *MySQLDatabase {
	// db, err := sql.Open("mysql", "user:password@/dbname")
	return &MySQLDatabase{
		uploads: make(map[uint]*domain.Upload),
	}
}

func (db *MySQLDatabase) GetUploadByUploadId(uploadId string) *domain.Upload {
	return nil
}

func (db *MySQLDatabase) CreateUpload(upload *domain.Upload) (uint, error) {
	return 0, nil
}
