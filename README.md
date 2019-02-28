# skynet
## Requirements
* Go
* PostgreSQL
## Getting started
### Prepare database
1. Create a new `skynet_dev` database
2. Export database connection environment variables, e.g.
```
export SKYNET_DEV_DB_HOST=localhost
export SKYNET_DEV_DB_PORT=5432
export SKYNET_DEV_DB_DBNAME=skynet_dev
export SKYNET_DEV_DB_USER=postgres
export SKYNET_DEV_DB_PASSWORD=postgres
```
### Run app
1. `cd skynet`
2. Install all dependencies with `go get ./...`
3. Start app with `go run cmd/skynet/main.go `
4. Navigate your browser to `http://localhost:3030/graphiql/`
5. Query data, e.g.
```
{
  prototypes {
    name
    units {
    	health
    }
  }
}
```
