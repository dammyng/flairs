package persistence

import (
	"auth/libs/utils"
	"log"

	"shared/models/appuser"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type MysqlLayer struct {
	Session *gorm.DB
}

//NewMysqlLayer - Application database layer
func NewMysqlLayer(dbconfig string) *MysqlLayer {
	s, err := gorm.Open("mysql", dbconfig)
	if err != nil {
		log.Fatalf("[createDBSession]: %s\n", err)
	}
	return &MysqlLayer{
		Session: s,
	}
}

// AddUser uses gorm's save method to add a new user to db
func (sqlLayer *MysqlLayer) AddUser(user appuser.User) error {
	return sqlLayer.Session.Create(utils.GRPCModelToSQL(&user)).Error
}

func (sqlLayer *MysqlLayer) AllUsers() ([]appuser.User, error) {
	users := []appuser.User{}
	result := sqlLayer.Session.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
func (sqlLayer *MysqlLayer) FindUsers() ([]appuser.User, error) {
	users := []appuser.User{}
	return users, nil
}

func (sqlLayer *MysqlLayer) UpdateUser(data *appuser.UpdateArg) error {
	return sqlLayer.Session.Model(&data.OldUser).Updates(data.NewObj).Error
}

func (sqlLayer *MysqlLayer) GetUser(in *appuser.User) (appuser.User, error) {
	var user appuser.User
	if sqlLayer.Session.Where(in).First(&user).RecordNotFound() {
		return user, gorm.ErrRecordNotFound
	}
	return user, nil
}

func (db *MysqlLayer) GetFreshSession() *gorm.DB {
	return db.Session.New()
}

func (db *MysqlLayer) DoMigrations() {
	//session := db.GetFreshSession()
	db.Session.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&appuser.User{})
}

func (db *MysqlLayer) Close() {
	db.Session.Close()
}
