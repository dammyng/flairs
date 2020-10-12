package v1

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/gomodule/redigo/redis"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1internals "auth/internals/v1"
	redisconn "auth/redis"

	"auth/libs/setup"
	"auth/pkg/api/v1"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var testDb *gorm.DB
var testRedis redis.Conn

func TestMain(m *testing.M) {

	initDB()
	initRedis()
	code := m.Run()
	//clearDB()
	os.Exit(code)
}
func initDB() {
	var err error
	dsn := fmt.Sprintf("root:password@tcp(127.0.0.1)/alpha_plus?charset=utf8&parseTime=True&loc=Local")
	testDb, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Panicf("flairServiceServer.AddNewUser() error = %v,", err)
	}
	testDb.Exec(setup.SQLMode)
}

func initRedis() {
	redisPool := redisconn.NewPool("localhost:6379")
	testRedis = redisPool.Get()
}

func TestAddUser_ok(t *testing.T) {
	clearUsersTable()
	ctx := context.Background()
	sqlLayer := v1internals.NewMysqlLayer(testDb)
	s := NewFlairsServiceServer(sqlLayer, testRedis)

	req := &v1.AddNewUserRequest{
		Api:   "v1",
		Email: "el.s",
		Ref:   "dddddd",
	}

	got, err := s.AddNewUser(ctx, req)

	if err != nil{
		t.Errorf("flairServiceServer.AddNewUser() error = %v, wantErr %v", err, "f")
		return
	}

	var u v1internals.User
	testDb.Where("email = ?", req.Email).Last(&u)

	if err == nil && !reflect.DeepEqual(got.ID, u.ID) {
		t.Errorf("flairServiceServer.AddNewUser() = %v, want %v", got.ID, u.ID)
	}
}

func TestAddUser_duplicate_email(t *testing.T) {
	clearUsersTable()
	ctx := context.Background()
	sqlLayer := v1internals.NewMysqlLayer(testDb)
	s := NewFlairsServiceServer(sqlLayer, testRedis)

	req := &v1.AddNewUserRequest{
		Api:   "v1",
		Email: "el.s",
		Ref:   "dddddd",
	}

	_, err := s.AddNewUser(ctx, req)
	_, duplicateErr := s.AddNewUser(ctx, req)

	if err != nil{
		t.Errorf("flairServiceServer.AddNewUser() error = %v, wantErr %v", err, "f")
		return
	}

	if duplicateErr == nil{
		t.Errorf("flairServiceServer.AddNewUser() should have returned a duplicate error  = %v, wantErr %v", err, "f")
		return
	}

	if (duplicateErr != nil) && duplicateErr.Error() != status.Error(codes.AlreadyExists, "A Flairs account with this email already exists. Please try logging in.").Error(){
		t.Errorf("flairServiceServer.AddNewUser() should return a duplicate error but got = %v,", duplicateErr.Error())
	}
}


func TestAddUser_invalid_entry(t *testing.T) {
	clearUsersTable()
	ctx := context.Background()
	sqlLayer := v1internals.NewMysqlLayer(testDb)
	s := NewFlairsServiceServer(sqlLayer, testRedis)

	req := &v1.AddNewUserRequest{
		Api:   "v1",
		Ref:   "dddddd",
	}

	_, err := s.AddNewUser(ctx, req)

	if err == nil{
		t.Errorf("flairServiceServer.AddNewUser() should have returned an internal server error  = %v", err)
		return
	}

	if (err != nil) && err.Error() != status.Error(codes.Internal, "failed to insert into Users-> "+err.Error()).Error(){
		t.Errorf("flairServiceServer.AddNewUser() Wrong error! should return a internal error but got = %v,", err.Error())
	}
}

func clearUsersTable() {
	testDb.Exec(setup.ClearUserTable)
}