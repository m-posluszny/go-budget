package auth_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/google/go-querystring/query"
	"github.com/m-posluszny/go-ynab/src/auth"
	"github.com/m-posluszny/go-ynab/src/db"
	"github.com/m-posluszny/go-ynab/src/misc_tests"
)

func TestSessionRegister(t *testing.T) {
	_, mockw := db.InitMockDbs()
	misc_tests.MockHashPassword()
	form := MockRegisterForm()
	w, s := misc_tests.GetGinMock()

	mockCreds := MockCredentials()
	mockCreds.PasswordHash = form.DbView().PasswordHash
	MockCreateUser(&mockw, mockCreds)
	MockGetByUsername(&mockw, mockCreds)
	formValues, _ := query.Values(form)
	req, _ := http.NewRequest("POST", "/register", strings.NewReader(formValues.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	s.ServeHTTP(w, req)

	if code := w.Code; code != 302 {
		t.Error("Wrong status code", code, formValues.Encode())

	}
}

func MockLogin(w *httptest.ResponseRecorder, s *gin.Engine, mockr sqlmock.Sqlmock, form auth.LoginForm) string {
	mockCreds := MockCredentials()
	mockCreds.PasswordHash = form.DbView().PasswordHash
	MockGetByUsername(&mockr, mockCreds)
	MockGetByUsername(&mockr, mockCreds)
	formValues, _ := query.Values(form)
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(formValues.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	s.ServeHTTP(w, req)
	return formValues.Encode()
}

func TestSessionLogin(t *testing.T) {
	w, s, mockr, _ := misc_tests.MockInitMockServer()
	formValues := MockLogin(w, s, mockr, MockLoginForm())
	if code := w.Code; code != 302 {
		t.Error("Wrong status code", code, formValues)
	}
}

func TestSessionLogout(t *testing.T) {
	w, s, mockr, _ := misc_tests.MockInitMockServer()
	MockLogin(w, s, mockr, MockLoginForm())

	req, _ := http.NewRequest("GET", "/logout", nil)
	s.ServeHTTP(w, req)
	if code := w.Code; code != 302 {
		t.Error("Wrong status code", code)

	}
	req, _ = http.NewRequest("GET", "/panel/", nil)
	s.ServeHTTP(w, req)
	if code := w.Code; code != 302 {
		t.Error("Wrong status code", code)
	}
	if !misc_tests.CompareBody(w, `<a href="/login">Found`) {
		t.Error("Wrong redirection", w.Body.String())
	}

}

func TestAccessPanel(t *testing.T) {
	w, s, mockr, _ := misc_tests.MockInitMockServer()
	MockLogin(w, s, mockr, MockLoginForm())
	req, _ := http.NewRequest("GET", "/panel/", nil)
	req.AddCookie(w.Result().Cookies()[0])
	creds := MockCredentials()
	MockGetByUid(&mockr, creds)
	MockGetByUid(&mockr, creds)
	s.ServeHTTP(w, req)
	if code := w.Code; code != 302 {
		t.Error("Wrong status code", code)
	}
	if !misc_tests.CompareBody(w, `<title>GoBudget - Panel</title>`) {
		t.Error("Wrong redirection")
	}

}

func TestAccessPanelFail(t *testing.T) {
	w, s, _, _ := misc_tests.MockInitMockServer()
	req, _ := http.NewRequest("GET", "/panel/", nil)
	creds := MockCredentials()
	creds.Uid = "-"
	s.ServeHTTP(w, req)
	if code := w.Code; code != 302 {
		t.Error("Wrong status code", code)
	}
	if !misc_tests.CompareBody(w, `<a href="/login">Found`) {
		t.Error("Wrong redirection", w.Body.String())
	}

}
