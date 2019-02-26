package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jinzhu/gorm"

	"github.com/lmuench/skynet/internal/app/skynet/middleware"
	"github.com/lmuench/skynet/internal/platform/orm"
	"github.com/urfave/negroni"
)

var n *negroni.Negroni
var db *gorm.DB

func TestMain(m *testing.M) {
	db = orm.InitPostgresTest()
	defer db.Close()

	n = negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		// negroni.NewStatic(http.Dir("web/static")),
	)
	n.UseFunc(middleware.CORS)

	code := m.Run()

	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	n.ServeHTTP(rec, req)
	return rec
}
