package persistence

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type MysqlLayer struct {
	Session *gorm.DB
}

func NewMysqlLayer(dbconfig DBConfig) *MysqlLayer {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s/?charset=utf8&parseTime=True&loc=Local", dbconfig.Username, dbconfig.Password, dbconfig.Hosts, dbconfig.Database)
	s,err := gorm.Open("mysql", dsn)
	if err !=nil{
		log.Fatalf("[createDBSession]: %s\n", err)
	}
	return &MysqlLayer{
		Session: s,
	}
}