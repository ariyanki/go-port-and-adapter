
# Port & Adapter

install https://github.com/swaggo/swag

Run Service:
1. Copy config_example.json to config.json
2. Run MySQL DB Migration for mac M1, go to folder migrations/apple-silicon, run `./gg-mysql-migrate up` or build migration for another platform with https://github.com/abcdef-id/go-migrations
3. Run Api, run `./run_dev.sh`, or
4. Run Manually:
	- Update swagger, run `swag init -g apps/http/main.go`
	- Run API, run `go run apps/http/main.go`
5. API Url, {hostname}:{port}/api/v1/{module}