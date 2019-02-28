package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/lmuench/skynet/internal/app/skynet/middleware"
	"github.com/lmuench/skynet/internal/pkg/services"
	"github.com/lmuench/skynet/internal/pkg/services/types"
	"github.com/lmuench/skynet/internal/platform/orm"
	"github.com/samsarahq/thunder/graphql"
	"github.com/samsarahq/thunder/graphql/graphiql"
	"github.com/samsarahq/thunder/graphql/introspection"
	"github.com/samsarahq/thunder/graphql/schemabuilder"
	"github.com/urfave/negroni"
)

func main() {
	// db
	db, adm := orm.InitPostgresDev()
	defer db.Close()

	// test data
	for i := 1; i <= 10; i++ {
		num := strconv.Itoa(i)
		pt := types.Prototype{
			Name:  "T-" + num + "000",
			Units: []types.Unit{},
		}
		for j := 0; j < 10; j++ {
			u := types.Unit{Health: 100}
			pt.Units = append(pt.Units, u)
		}
		db.Create(&pt)
	}

	// services
	units := services.Units{DB: db}
	prototypes := services.Prototypes{DB: db}

	// graphql schema
	builder := schemabuilder.NewSchema()
	builder.Query().FieldFunc("units", units.GetUnits)
	builder.Object("Unit", types.Unit{}).FieldFunc("createdAt", func(u *types.Unit) string {
		return u.CreatedAt.String()
	})
	builder.Query().FieldFunc("unit", func(args struct{ ID string }) types.Unit {
		id, err := strconv.Atoi(args.ID)
		if err != nil {
			return types.Unit{}
		}
		return units.GetUnit(id)
	})
	builder.Query().FieldFunc("prototypes", prototypes.GetPrototypes)
	builder.Query().FieldFunc("prototype", func(args struct{ ID string }) types.Prototype {
		id, err := strconv.Atoi(args.ID)
		if err != nil {
			return types.Prototype{}
		}
		return prototypes.GetPrototype(id)
	})
	builder.Mutation().FieldFunc("noop", func() {}) // needs a mutation, else auto-complete won't work
	schema := builder.MustBuild()
	introspection.AddIntrospectionToSchema(schema)

	// middleware
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
	)
	n.UseFunc(middleware.CORS)

	// mux
	mux := http.NewServeMux()
	mux.Handle("/graphql", graphql.Handler(schema))
	mux.Handle("/graphiql/", http.StripPrefix("/graphiql/", graphiql.Handler()))
	adm.MountTo("/admin", mux)
	n.UseHandler(mux)

	// http server
	log.Println("Server listening on port 3030")
	log.Fatal(http.ListenAndServe(":3030", n))
}
