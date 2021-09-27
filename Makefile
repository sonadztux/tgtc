# run backend service
go_run_backend:
	@go build -v -o bin/backend_service backend/main.go
	@./bin/backend_service