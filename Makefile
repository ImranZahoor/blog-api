build:
	go build cmd/main.go

run-dev:
	nodemon --watch "./**/*.go" --exec "go" run cmd/main.go

run:
	go run cmd/main.go
