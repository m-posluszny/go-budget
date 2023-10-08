package auth

import (
	"fmt"

	"github.com/m-posluszny/go-ynab/src/db"
	"golang.org/x/crypto/bcrypt"
)

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
type RegisterForm struct {
	LoginForm
	RePassword string `form:"repassword" binding:"required"`
}

type Credentials struct {
	Uid          string
	Username     string
	PasswordHash []byte `db:"password_hash"`
}

const createUserQuery = "INSERT INTO credentials (username, uid, password_hash) VALUES (:username, gen_random_uuid(), :password_hash);"

func GetUserFromUid(dbx *db.DBRead, uid string) (*Credentials, error) {
	creds := Credentials{}
	err := dbx.Get(&creds, "SELECT uid, username, password_hash FROM credentials WHERE uid=$1;", uid)
	return &creds, err
}

func GetUserFromName(dbx *db.DBRead, username string) (*Credentials, error) {
	creds := Credentials{}
	err := dbx.Get(&creds, "SELECT uid, username, password_hash FROM credentials WHERE username=$1;", username)
	return &creds, err
}

func (form LoginForm) HashedPassword() []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return hashedPassword
}

func (form LoginForm) dbView() Credentials {
	return Credentials{Username: form.Username, PasswordHash: form.HashedPassword()}
}

func MatchPassword(dbx *db.DBRead, form LoginForm) bool {
	creds, err := GetUserFromName(dbx, form.Username)
	if err != nil {
		panic(err)
	}
	return bcrypt.CompareHashAndPassword(creds.PasswordHash, []byte(form.Password)) == nil
}

func CreateUser(form RegisterForm) (*Credentials, error) {
	dbx := db.GetDbWrite()
	tx := dbx.MustBegin()
	newUser := form.dbView()
	response, err := tx.NamedExec(createUserQuery, newUser)
	fmt.Println(response, err)
	tx.Commit()
	if err != nil {
		return nil, err
	}
	return GetUserFromName(dbx, newUser.Username)
}
