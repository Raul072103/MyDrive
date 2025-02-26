package main

import (
	"MyDrive/internal/auth"
	"MyDrive/internal/db"
	"MyDrive/internal/env"
	repository "MyDrive/internal/repo"
	"database/sql"
	"expvar"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"runtime"
	"time"
)

const version = "0.0.0"

//	@title			MyDrive API
//	@description	API for MyDrive.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath					/v1
//
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description
func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost:5434/mydrive?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env:         env.GetString("DEV", "development"),
		apiURL:      env.GetString("EXTERNAL_URL", "localhost:8080"),
		frontendURL: env.GetString("FRONTEND_URL", "http://localhost:5174"),
		auth: authConfig{tokenConfig{
			secret: env.GetString("AUTH_TOKEN_SECRET", "secret_example"),
			exp:    time.Hour * 24 * 3,
			iss:    "mydrive",
		}},
		drive: driveConfig{
			root: env.GetString("FILE_SYSTEM_ROOT_FOLDER", "./testing/root/"),
		},
	}

	// Logger
	logger := zap.Must(zap.NewDevelopment()).Sugar()
	defer func(logger *zap.SugaredLogger) {
		if err := logger.Sync(); err != nil {
			logger.Fatalf("failed cleaning up zap logger: %v", err)
		}
	}(logger)

	// Database
	postgresDB, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime)

	if err != nil {
		logger.Fatalf("failed connecting to the postgres DB: %v", err)
	}

	defer func(database *sql.DB) {
		if err := postgresDB.Close(); err != nil {
			logger.Fatalf("failed cleaning up zap logger: %v", err)
		}
	}(postgresDB)

	logger.Info("database connection pool established")
	logger.Infof("drive root folder: %s", cfg.drive.root)

	// Repository
	repo := repository.NewRepo(postgresDB)

	// Authenticator
	jwtAuthenticator := auth.NewJWTAuthenticator(
		cfg.auth.token.secret,
		cfg.auth.token.iss,
		cfg.auth.token.iss)

	app := &application{
		config:        cfg,
		repo:          repo,
		logger:        logger,
		authenticator: jwtAuthenticator,
	}

	// Metrics collected
	expvar.NewString("version").Set(version)
	expvar.Publish("database", expvar.Func(func() any {
		return postgresDB.Stats()
	}))
	expvar.Publish("goroutines", expvar.Func(func() any {
		return runtime.NumGoroutine()
	}))

	mux := app.mount()
	logger.Fatal(app.run(mux))
}
