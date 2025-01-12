.PHONY: install build dev tunnel migrate generate

install:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

build:
	skaffold build --file-output output.json

dev:
	skaffold dev -v debug -p dev --wait-for-deletions=true --kube-context=minikube

tunnel:
	#minikube service mariadb --url -n mariadb &
	minikube service mii-go-web --url -n mii-go

migrate:
	migrate -path cmd/app/migrations/mysql -database "mysql://root:password#123@tcp(localhost:56836)/mii" up

generate:
	sqlc generate