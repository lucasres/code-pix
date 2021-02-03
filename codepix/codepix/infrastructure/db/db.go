package db

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/lucasres/code-pix/domain/model"
	_ "gorm.io/driver/sqlite"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "../../.env")

	if err != nil {
		log.Fatal("Fail load .env in DB init", err)
	}
}

func CreateDB(env string) *gorm.DB {
	var db *gorm.DB
	var err error
	var dialect string
	var dsn string

	if env != "test" {
		dialect = os.Getenv("DB_TYPE")
		dsn = os.Getenv("DSN")
	} else {
		dialect = os.Getenv("DB_TYPE_TEST")
		dsn = os.Getenv("DSN_TEST")
	}

	db, err = gorm.Open(dialect, dsn)

	if err != nil {
		log.Fatal("Erro in open connection to database", err)
	}

	if os.Getenv("DEBUG") == "true" {
		db.LogMode(true)
	}

	if os.Getenv("AUTO_MIGRATE_DB") == "true" {
		db.AutoMigrate(
			&model.Bank{},
			&model.Account{},
			&model.PixKey{},
			&model.Transaction{},
		)
	}

	return db

}
