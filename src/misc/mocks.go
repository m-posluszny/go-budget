package misc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/m-posluszny/go-ynab/src/auth"
	"github.com/m-posluszny/go-ynab/src/config"
	"github.com/m-posluszny/go-ynab/src/db"
	"github.com/m-posluszny/go-ynab/src/server"
	"github.com/m-posluszny/go-ynab/src/stores"
)

func MockInitMockServer() (*httptest.ResponseRecorder, *gin.Engine, sqlmock.Sqlmock, sqlmock.Sqlmock) {
	mockr, mockw := db.InitMockDbs()
	MockHashPassword()
	MockMatchPassword(nil)
	w, s := GetGinMock()
	return w, s, mockr, mockw

}

func CompareBody(w *httptest.ResponseRecorder, body string) bool {
	b, err := regexp.MatchString(body, w.Body.String())
	if !b {
		fmt.Println("Looked for", body)
		fmt.Println("Got", w.Body.String())
		fmt.Println(b, err)
	}
	return b
}

func MockHashPassword() {
	auth.GenerateHashPassword = func(p []byte) ([]byte, error) { return make([]byte, 0), nil }
}

func MockMatchPassword(e error) {
	auth.CompareHashAndPassword = func(h, p []byte) error { return e }
}
func GetGinMock() (*httptest.ResponseRecorder, *gin.Engine) {
	gin.SetMode(gin.TestMode)
	srv := server.Init(config.Config{}, stores.GetMockSessionStore(config.AuthConf{"secret", 420}), "../templates/**/*")
	rec := httptest.NewRecorder()

	return rec, srv
}

func FetchExpects(t *testing.T, mock sqlmock.Sqlmock) {
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func GetJsonBuff(v any) *bytes.Buffer {
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(v)
	return payloadBuf
}
