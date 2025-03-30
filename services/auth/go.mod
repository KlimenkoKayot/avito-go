module github.com/klimenkokayot/avito-go/services/auth

go 1.23.7

require (
	github.com/google/uuid v1.6.0
	github.com/joho/godotenv v1.5.1
	github.com/klimenkokayot/avito-go/libs/logger v0.0.0-20250328223537-07b05408642a
	github.com/lib/pq v1.10.9
	golang.org/x/crypto v0.36.0
)

require github.com/gorilla/mux v1.8.1 // indirect

require (
	github.com/jmoiron/sqlx v1.4.0
	github.com/klimenkokayot/avito-go/libs/router v0.0.0-20250329212844-f3cd09b36437
	github.com/sirupsen/logrus v1.9.3 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
)
