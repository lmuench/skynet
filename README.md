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
1. `git clone https://github.com/lmuench/skynet.git`
2. `cd skynet`
3. Install all dependencies with `go get ./...`
4. Start app with `go run cmd/skynet/main.go `
5. Navigate your browser to http://localhost:3030/graphiql/
6. Query data, e.g.
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
### Admin Interface
http://localhost:3030/admin
