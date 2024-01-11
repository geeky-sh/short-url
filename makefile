swag:
	rm ./swagger.yaml
	swagger generate spec -w cmd/server -o ./swagger.yaml
	swagger serve -F=swagger swagger.yaml

liteup:
	~/go/bin/migrate -database sqlite3://cmd/server/main.db\?mode=rwc --path db/migrations up

litedown:
	~/go/bin/migrate -database sqlite3://cmd/server/main.db\?mode=rwc --path db/migrations down

dbdown:
	migrate -database postgres://aash:@localhost:5432/shorturl\?sslmode=disable --path db/migrations down

dbup:
	migrate -database postgres://aash:@localhost:5432/shorturl\?sslmode=disable --path db/migrations up

# example: make dbfix v=xxxxx
dbfix:
	migrate -database postgres://aash:@localhost:5432/shorturl\?sslmode=disable -path db/migrations force $v

# example: make dbcreate name=create_temps
dbcreate:
	migrate create -ext sql -dir db/migrations $(name)
