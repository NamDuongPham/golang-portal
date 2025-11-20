package config

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

var (
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	Cfg    = Load()
)

type Config struct {
	API        APIConfig `env:",prefix=API_"`
	Environemt string    `env:"ENV"`
	Databases  DBConfig  `env:",prefix=DB_"`
	JWT        JWTConfig `env:",prefix=JWT_"`
}

type JWTConfig struct {
	AccessSecretKey  string `env:"ACCESS_TOKEN_SECRET, default=your_access_secret_key"`
	RefreshSecretKey string `env:"REFRESH_TOKEN_SECRET, default=your_refresh_secret_key"`
}

type APIConfig struct {
	Host     string `env:"SERVER_HOST, default=0.0.0.0"`
	Port     int    `env:"SERVER_PORT, default=8000"`
	LogLevel string `env:"LOG_LEVEL, default=WARN"`
}

type DBConfig struct {
	UseReplication  bool          `env:"USE_REPLICATION, default=false"`
	Source          DBAttributes  `env:",prefix=SOURCE_"`
	Replica         DBAttributes  `env:",prefix=REPLICA_"`
	LogLevel        string        `env:"LOG_LEVEL, default=WARN"`
	ConnMaxIdleTime time.Duration `env:"CONN_MAX_IDLE_TIME, default=300s"`
	ConnMaxLifeTime time.Duration `env:"CONN_MAX_LIFE_TIME, default=300s"`
	MaxIdleConns    int           `env:"MAX_IDLE_CONNS, default=5"`
	MaxOpenConns    int           `env:"MAX_OPEN_CONNS, default=10"`
}

type DBAttributes struct {
	Dsn      string `env:"DSN"`
	Host     string `env:"HOST, default=127.0.0.1"`
	Port     int    `env:"PORT, default=5432"`
	Database string `env:"DATABASE, default=mydb"`
	Username string `env:"USERNAME, default=admin"`
	Password string `env:"PASSWORD, default=admin123"`
	SSLMode  string `env:"SSL_MODE, default=disable"`
}

func (c DBAttributes) GetPostgresDSN() string {
	if c.Dsn != "" {
		return c.Dsn
	}
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s", c.Username, c.Password, c.Database, c.Host, c.Port, c.SSLMode)
}

func Load() Config {
	ctx := context.Background()

	// Try to load a .env file (optional). If it doesn't exist, continue —
	// go-envconfig will read environment variables from the environment.
	if err := godotenv.Load(); err != nil {
		// Not fatal; often there is no .env in production. Log at debug level.
		logger.LogAttrs(ctx, slog.LevelDebug, "No .env file loaded", slog.Any("error", err))
	}

	var c Config
	if err := envconfig.Process(ctx, &c); err != nil {
		logger.ErrorContext(ctx, "Failed to load config", slog.Any("error", err))
	}
	return c
}
