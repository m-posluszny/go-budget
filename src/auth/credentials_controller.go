package auth

import (
	"github.com/m-posluszny/go-ynab/src/db"
	"golang.org/x/crypto/bcrypt"
)

type LoginForm struct {
	Username string `form:"username" binding:"required" url:"username"`
	Password string `form:"password" binding:"required" url:"password"`
}
type RegisterForm struct {
	LoginForm
	RePassword string `form:"repassword" binding:"required" url:"repassword"`
}

type Credentials struct {
	Uid          string
	Username     string
	PasswordHash []byte `db:"password_hash"`
}

var GenerateHashPassword = func(p []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(p, bcrypt.DefaultCost)
}

var CompareHashAndPassword = func(h []byte, p []byte) error {
	return bcrypt.CompareHashAndPassword(h, p)

}

func (form LoginForm) HashedPassword() []byte {
	hashedPassword, err := GenerateHashPassword([]byte(form.Password))
	if err != nil {
		panic(err)
	}
	return hashedPassword
}

func (form LoginForm) DbView() Credentials {
	return Credentials{Username: form.Username, PasswordHash: form.HashedPassword()}
}

func GetUserFromUid(dbx *db.DBRead, uid string) (*Credentials, error) {
	creds := Credentials{}
	err := dbx.Get(&creds,
		`SELECT uid, username, password_hash FROM credentials WHERE uid=$1;`,
		uid)
	return &creds, err
}

func GetUserFromName(dbx *db.DBRead, username string) (*Credentials, error) {
	creds := Credentials{}
	err := dbx.Get(&creds,
		`SELECT uid, username, password_hash FROM credentials WHERE username=$1;`,
		username)
	return &creds, err
}
func CreateUser(dbx *db.DBWrite, newUser Credentials) (*Credentials, error) {
	_, err := dbx.NamedExec(
		`INSERT INTO credentials (username, uid, password_hash) VALUES (:username, gen_random_uuid(), :password_hash);`,
		newUser)
	if err != nil {
		return nil, err
	}
	return GetUserFromName(dbx, newUser.Username)
}

func MustMatchPassword(dbx *db.DBRead, form LoginForm) bool {
	creds, err := GetUserFromName(dbx, form.Username)
	if err != nil {
		panic(err)
	}
	return CompareHashAndPassword(creds.PasswordHash, []byte(form.Password)) == nil
}
