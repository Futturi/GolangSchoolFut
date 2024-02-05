package repositoryinitf

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Port     string
	Hostname string
	Username string
	Password string
	NameDB   string
	Sslmode  string
}

func InitPostgre(cfg Config) (*sqlx.DB, error) {
	con, err := sqlx.Connect("postgres", fmt.Sprintf("host =%s port =%s user =%s dbname=%s password=%s sslmode=%s",
		cfg.Hostname, cfg.Port, cfg.Username, cfg.NameDB, cfg.Password, cfg.Sslmode))
	if err != nil {
		return nil, err
	}
	return con, nil
}
