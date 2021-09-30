# run backend service
go_run_backend:
	@go build -v -o bin/backend_service backend/main.go
	@./bin/backend_service

# run gql service
go_run_gql:
	@go build -v -o bin/gql_service gql/main.go
	@./bin/gql_service
