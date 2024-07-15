package db

import (
	"database/sql"
	"fmt"

	"github.com/pleimer/ticketer/server/ent"
	"go.uber.org/zap"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DBConnectionConfig struct {
	Host     string `short:"h" long:"host" env:"TICKETER_DB_HOST" default:"localhost" description:"Database host"`
	Port     int    `short:"p" long:"port" env:"TICKETER_DB_PORT" default:"5432" description:"Database port"`
	User     string `short:"u" long:"user" env:"TICKETER_DB_USER" required:"true" description:"Database user"`
	DBName   string `short:"d" long:"dbname" env:"TICKETER_DB_NAME" default:"ticketerdb" description:"Database name"`
	Password string `long:"password" env:"TICKETER_DB_PASSWORD" required:"true" description:"Database password"`
	SSLMode  string `long:"sslmode" env:"TICKETER_DB_SSLMODE" default:"disable" description:"SSL mode (disable, require, verify-ca, verify-full)"`
}

// PGConnectionURL returns a formatted postgres connection URL
func (c DBConnectionConfig) PGConnectionURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c.User, c.Password, c.Host, c.Port, c.DBName, c.SSLMode)
}

type DB struct {
	logger *zap.Logger
	db     *sql.DB
	Client *ent.Client
}

func NewDB(logger *zap.Logger) *DB {
	return &DB{
		logger: logger,
	}
}

func (d *DB) Open(cfg DBConnectionConfig) {
	var err error

	d.db, err = sql.Open("pgx", cfg.PGConnectionURL())
	if err != nil {
		d.logger.Fatal("connecting to db", zap.Error(err))
	}

	driver := entsql.OpenDB(dialect.Postgres, d.db)
	d.Client = ent.NewClient(ent.Driver(driver))
}

func (d *DB) Close() {
	err := d.db.Close()
	if err != nil {
		d.logger.Error("closing db connection", zap.Error(err))
	}
}
