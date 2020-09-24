package persistence

import (
	"fmt"
	"log"
	
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"shared/models/appuser"
)

type MysqlLayer struct {
	Session *gorm.DB
}

//NewMysqlLayer - Application database layer
func NewMysqlLayer(dbconfig DBConfig) *MysqlLayer {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/?charset=utf8&parseTime=True&loc=Local", dbconfig.Username, dbconfig.Password, dbconfig.Hosts)
	s,err := gorm.Open("mysql", dsn)
	if err !=nil{
		log.Fatalf("[createDBSession]: %s\n", err)
	}
	return &MysqlLayer{
		Session: s,
	}
}
// AddUser uses gorm's save method to add a new user to db
func (sqlLayer *MysqlLayer) AddUser(user appuser.User)error  {
	return nil
}

func (sqlLayer *MysqlLayer) AllUsers()([]appuser.User, error)  {
	users := []appuser.User{}
	result := sqlLayer.Session.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}