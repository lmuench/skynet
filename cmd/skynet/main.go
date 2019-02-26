package main

import (
	"log"
	"net/http"

	"github.com/lmuench/skynet/internal/app/skynet/middleware"
	"github.com/lmuench/skynet/internal/platform/orm"
	"github.com/urfave/negroni"
)

func main() {
	db, adm := orm.InitPostgresDev()
	defer db.Close()

	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		// negroni.NewStatic(http.Dir("web/static")),
	)
	n.UseFunc(middleware.CORS)

	mux := http.NewServeMux()
	// mux.Handle("/graphql", )
	// mux.Handle("/graphiql", )
	adm.MountTo("/admin", mux)
	n.UseHandler(mux)

	log.Println("Server listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", n))
}
