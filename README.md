# Port & Adapter

1. Copy config_example.json to config.json
2. Run MySQL DB Migration for mac M1, go to folder migrations/apple-silicon, run `./gg-mysql-migrate up`
3. Run Internal API, run `go run app/api/intl/main.go`
4. API Url, {hostname}:{port}/api/v1/{module}

![Image result for hexagonal architecture software](https://herbertograca.files.wordpress.com/2017/03/hexagonal-arch-4-ports-adapters2.png?w=708)