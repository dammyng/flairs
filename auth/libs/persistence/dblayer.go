package persistence

import (
	"auth/libs/setup"
	"auth/libs/utils"
	"log"
	"os"

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
// FindUsers - return a list of users 
func (sqlLayer *MysqlLayer) FindUsers() ([]appuser.User, error) {
	users := []appuser.User{}
	return users, nil
}
// UpdateUser - update a user but not in the DB
func (sqlLayer *MysqlLayer) UpdateUser(data *appuser.UpdateArg) error {
	session := sqlLayer.GetFreshSession()
	return session.Model(&data.OldUser).Updates(&data.NewObj).Error
}

// GetUser - return a sure from the database
func (sqlLayer *MysqlLayer) GetUser(in *appuser.User) (*appuser.User, error) {
	session := sqlLayer.GetFreshSession()
	var user appuser.User
	if session.Where(in).First(&user).RecordNotFound() {
		log.Printf("not found")
		return nil, gorm.ErrRecordNotFound
	}
	return &user, nil
}

func (db *MysqlLayer) GetFreshSession() *gorm.DB {
	session := db.Session.New()
	session.Exec(setup.SetTimeZone)
	session.Exec(setup.UseAlphaPlus)
	return session
}

func (db *MysqlLayer) DoMigrations() {
	s := NewMysqlLayer(os.Getenv("DBConnString"))
	s.Session.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(appuser.User{})
}

func (db *MysqlLayer) Close() {
	db.Session.Close()
}
