package database

import (
	"Adeeb_Go/internal/config"
	"Adeeb_Go/internal/database/sqlc"
	"Adeeb_Go/logger"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() (*pgxpool.Pool, error) {
	ctx := context.Background()
	// Create database connection
	connPool, err := pgxpool.NewWithConfig(ctx, GetPoolConfig())
	if err != nil {
		logger.Error().Stack().Err(err).Msg("Database not connected")
		os.Exit(1)
	}
	logger.Info().Msg("Database Connected")

	c, err := connPool.Acquire(ctx)
	defer c.Release()

	dataTypeNames := []string{ // An underscore prefix is an array type in pgtypes.
		"role",
		"_role",
		"time_period",
		"_time_period",
		"verse",
		"_verse",
	}
	conn := c.Conn()
	for _, typeName := range dataTypeNames {
		dataType, err := conn.LoadType(ctx, typeName)
		if err != nil {
			logger.Error().Err(err).Msg("Couldb't register database type")
			os.Exit(1)
		}
		conn.TypeMap().RegisterType(dataType)
	}

	sqlc.Q = sqlc.New(connPool)

	return connPool, nil
}

func GetPoolConfig() *pgxpool.Config {
	connStr := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_NAME, config.DB_PASSWORD)

	dbConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		logger.Error().Stack().Err(err).Msg("Failed to create a config, error: ")
		os.Exit(1)
	}

	dbConfig.MaxConns = 4
	dbConfig.MinConns = 0
	dbConfig.MaxConnLifetime = time.Hour
	dbConfig.MaxConnIdleTime = time.Minute * 30
	dbConfig.HealthCheckPeriod = time.Minute
	dbConfig.ConnConfig.ConnectTimeout = time.Second * 5

	dbConfig.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
		logger.Info().Msg("Before acquiring the connection pool to the database!!")
		return true
	}

	dbConfig.AfterRelease = func(c *pgx.Conn) bool {
		logger.Info().Msg("After releasing the connection pool to the database!!")
		return true
	}

	dbConfig.BeforeClose = func(c *pgx.Conn) {
		logger.Info().Msg("Closed the connection pool to the database!!")
	}

	return dbConfig
}
