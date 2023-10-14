package auth_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/m-posluszny/go-ynab/src/auth"
	"github.com/m-posluszny/go-ynab/src/misc"
)

func MockValidateCredentials(creds auth.Credentials, mock auth.Credentials, t *testing.T) {
	if creds.Uid != mock.Uid {
		t.Error("Uid hash don't match")
	}
	if creds.Username != mock.Username {
		t.Error("Username hash don't match")
	}
	if !bytes.Equal(creds.PasswordHash, mock.PasswordHash) {
		t.Error("Password hash don't match")
	}
}

func TestGetUserFromUid(t *testing.T) {
	db, mock := misc.GetMockDb()

	mockCreds := auth.Credentials{Uid: "a-b-c-d", Username: "abcd", PasswordHash: []byte{1, 2, 3}}

	mock.ExpectQuery(
		`SELECT uid, username, password_hash FROM credentials WHERE uid=\$1;`).WithArgs("a-b-c-d").WillReturnRows(mock.NewRows([]string{"uid", "username", "password_hash"}).AddRow(mockCreds.Uid, mockCreds.Username, mockCreds.PasswordHash))

	creds, err := auth.GetUserFromUid(db, mockCreds.Uid)
	if err != nil {
		fmt.Println(err)
		t.Error("This test should not create err")
	}
	MockValidateCredentials(*creds, mockCreds, t)
	misc.FetchExpects(t, mock)

}

func TestGetUserFromUidErr(t *testing.T) {
	db, _ := misc.GetMockDb()
	_, err := auth.GetUserFromUid(db, "")
	if err == nil {
		fmt.Println(err)
		t.Error("This test should create err")
	}

}

func TestCreateUser(t *testing.T) {
	db, _ := misc.GetMockDb()
	_, err := auth.GetUserFromUid(db, "")
	if err == nil {
		fmt.Println(err)
		t.Error("This test should create err")
	}

}
