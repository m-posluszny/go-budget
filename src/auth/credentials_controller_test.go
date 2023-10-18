package auth_test

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/m-posluszny/go-ynab/src/auth"
	"github.com/m-posluszny/go-ynab/src/misc"
)

func MockLoginForm() auth.LoginForm {
	return auth.LoginForm{Username: "abcd", Password: "abcd1234"}
}

func MockCredentials() auth.Credentials {
	form := MockLoginForm()
	return auth.Credentials{Uid: "a-b-c-d", Username: form.Username, PasswordHash: make([]byte, 0)}
}

func CredsToRow(m sqlmock.Sqlmock, c auth.Credentials) *sqlmock.Rows {
	return m.NewRows([]string{"uid", "username", "password_hash"}).AddRow(c.Uid, c.Username, c.PasswordHash)

}

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

	mockCreds := MockCredentials()
	mock.ExpectQuery(
		`SELECT uid, username, password_hash FROM credentials WHERE uid=\$1;`).WithArgs(mockCreds.Uid).WillReturnRows(CredsToRow(mock, mockCreds))

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

func TestGetUserFromName(t *testing.T) {
	db, mock := misc.GetMockDb()

	mockCreds := MockCredentials()
	mock.ExpectQuery(
		`SELECT uid, username, password_hash FROM credentials WHERE username=\$1;`).WithArgs(mockCreds.Username).WillReturnRows(CredsToRow(mock, mockCreds))

	creds, err := auth.GetUserFromName(db, mockCreds.Username)
	if err != nil {
		fmt.Println(err)
		t.Error("This test should not create err")
	}
	MockValidateCredentials(*creds, mockCreds, t)
	misc.FetchExpects(t, mock)

}

func TestGetUserFromNameErr(t *testing.T) {
	db, _ := misc.GetMockDb()
	_, err := auth.GetUserFromName(db, "")
	if err == nil {
		fmt.Println(err)
		t.Error("This test should create err")
	}
}

func TestCreateUser(t *testing.T) {
	db, mock := misc.GetMockDb()
	mockCreds := MockCredentials()
	mock.ExpectExec(`INSERT INTO credentials \(username, uid, password_hash\) VALUES \(\?, gen_random_uuid\(\), \?\);`).WithArgs(mockCreds.Username, mockCreds.PasswordHash).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectQuery(
		`SELECT uid, username, password_hash FROM credentials WHERE username=\$1;`).WithArgs(mockCreds.Username).WillReturnRows(CredsToRow(mock, mockCreds))
	_, err := auth.CreateUser(db, mockCreds)
	if err != nil {
		fmt.Println(err)
		t.Error("This test should not create err")
	}
	misc.FetchExpects(t, mock)

}

func TestCreateUserErr(t *testing.T) {
	db, mock := misc.GetMockDb()
	mockCreds := MockCredentials()
	mock.ExpectExec(`INSERT INTO credentials \(username, uid, password_hash\) VALUES \(\?, gen_random_uuid\(\), \?\);`).WithArgs(mockCreds.Username, mockCreds.PasswordHash).WillReturnError(errors.New("db write error"))
	_, err := auth.CreateUser(db, mockCreds)
	if err == nil {
		fmt.Println(err)
		t.Error("This test should create err")
	}
	misc.FetchExpects(t, mock)

}
