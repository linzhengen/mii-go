.PHONY: install build dev tunnel

install:
	go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

build:
	skaffold build --file-output output.json

dev:
	skaffold dev -v debug -p dev --wait-for-deletions=true --kube-context=minikube

tunnel:
	minikube service mii-go --url -n ddd-gin-admin
