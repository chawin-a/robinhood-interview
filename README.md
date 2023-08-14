docker compose -f deployments/docker-compose.db.yml up -d
docker compose -f deployments\docker-compose.db.yml down

migrate -database postgres://postgres:password@localhost:5432/db?sslmode=disable -path migrations up
migrate -database postgres://postgres:password@localhost:5432/db?sslmode=disable -path migrations down -all

migrate create -ext sql -dir migrations -seq name_version

![alt text](https://github.com/chawin-a/robinhood-interview/blob/master/images/db-diagram.png?raw=true)