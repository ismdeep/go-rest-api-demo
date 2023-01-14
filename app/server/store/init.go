package store

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/ismdeep/go-rest-api-demo/app/server/conf"
	"github.com/ismdeep/go-rest-api-demo/internal/model"
)

var db *gorm.DB

func init() {
	switch conf.Basic.DB.Dialect {
	case "mysql", "mariadb":
		tmpDB, err := gorm.Open(mysql.Open(conf.Basic.DB.DSN))
		if err != nil {
			panic(err)
		}
		db = tmpDB
	case "sqlite":
		tmpDB, err := gorm.Open(sqlite.Open(fmt.Sprintf("%v/%v", conf.Basic.System.Data, conf.Basic.DB.DSN)))
		if err != nil {
			panic(err)
		}
		db = tmpDB
	default:
		panic(fmt.Errorf("unsupported db dialect. [%v]", conf.Basic.DB.Dialect))
	}

	if err := db.AutoMigrate(&model.User{}); err != nil {
		panic(err)
	}
}
