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

func (db *MysqlLayer) AddCardRequest(cardReq appuser.CardRequest) error {
	session := db.GetFreshSession()
	return session.Save(&cardReq).Error
}


func (db *MysqlLayer) AddWallet(wallet appuser.Wallet) error {
	session := db.GetFreshSession()
	return session.Save(&wallet).Error
}

func (db *MysqlLayer) FindUserCardRequests(id string) ([]appuser.CardRequest, error) {
	session := db.GetFreshSession()
	rows, err := session.Model(&appuser.CardRequest{}).Where("user_id = ?", id).Select("id,user_id,color,currency").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var _cr appuser.CardRequest
	var cr []appuser.CardRequest
	for rows.Next() {
		if err := rows.Scan(&_cr.ID, &_cr.UserId, &_cr.Color, &_cr.Currency); err != nil {
			log.Fatalln(err.Error())
		}
		//	log.Println(_cr.ID)
		//session.ScanRows(rows, &_cr)
		cr = append(cr, _cr)
	}
	//log.Println(len(cr))
	return cr, err
}

func (db *MysqlLayer) FindCardRequestById(id string) (appuser.CardRequest, error) {
	session := db.GetFreshSession()
	cardReq := appuser.CardRequest{}
	//fmt.Println("id is : ", id)
	if session.Model(&cardReq).Where("id = ? ", id).First(&cardReq).RecordNotFound() {
		return cardReq, gorm.ErrRecordNotFound
	}
	return cardReq, nil
}

func (db *MysqlLayer) FindUserWallets(id string) ([]appuser.Wallet, error) {
	session := db.GetFreshSession()
	rows, err := session.Model(&appuser.Wallet{}).Where("user_id = ?", id).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var _cr appuser.Wallet
	var cr []appuser.Wallet
	for rows.Next() {
		if err := rows.Scan(&_cr.WalletSig, &_cr.UserId, &_cr.WalletNo); err != nil {
			log.Fatalln(err.Error())
		}
		//	log.Println(_cr.ID)
		//session.ScanRows(rows, &_cr)
		cr = append(cr, _cr)
	}
	//log.Println(len(cr))
	return cr, err
}

func (db *MysqlLayer) FindWalletById(id string) (appuser.Wallet, error) {
	session := db.GetFreshSession()
	wallet := appuser.Wallet{}
	//fmt.Println("id is : ", id)
	if session.Model(&wallet).Where("id = ? ", id).First(&wallet).RecordNotFound() {
		return wallet, gorm.ErrRecordNotFound
	}
	return wallet, nil
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
