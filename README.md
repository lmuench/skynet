# skynet

## Getting started
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
