package conversordb

import (
	"github.com/YuLaFreitas/ConversoDB/configs"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ConversoDB struct {
	db *sqlx.DB
}

func NewMysqlRepository(cfg configs.MysqlDatabaseSetting) (*ConversoDB, error) {
	db, err := sqlx.Connect("pgx", cfg.String())
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(10000000000)
	db.SetMaxOpenConns(10)

	return NewMysqlRepositoryByBd(db)
}

func NewMysqlRepositoryByBd(db *sqlx.DB) (*ConversoDB, error) {
	return &ConversoDB{db: db}, nil
}
