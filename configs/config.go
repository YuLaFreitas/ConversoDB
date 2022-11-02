package configs

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

const (
	contextConfig = "config"
)

type AppConfig struct {
	//TODO: é preciso colocar requeride??
	Mysql               MysqlDatabaseSetting
	WorkdaysServiceHost string `env:"WORDKDAYS_SERVICE_HOST"`
	PathFIleMigrations  string `env:"PATH_FILES_MIGRATIONS"`
}

type MysqlDatabaseSetting struct {
	Host   string `env:"MYSQL_HOST,required"`
	Port   string `env:"MYSQL_PORT,required"`
	DbName string `env:"MYSQL_DB_NAME,required"`
	Schema string `env:"MYSQL_SCHEMA,required"`
	User   string `env:"MYSQL_USER,required"`
	Pwd    string `env:"MYSQL_PWD,required"`
}

func (ref MysqlDatabaseSetting) String() string {
	return fmt.Sprintf("mysql://%s:%s@%s:%s/%s?search_path=%s", ref.User, strings.Trim(ref.Pwd, "\n"), ref.Host, ref.Port, ref.DbName, ref.Schema)
}

func (ref MysqlDatabaseSetting) MySQLConnect() error {
	db, err := sql.Open("mysql", "root:"+ref.Pwd+"@tcp("+ref.Host+")/"+ref.DbName)

	if err != nil {
		return err
	}

	defer db.Close()
	return nil

}

func NewContextWithConfig(ctx context.Context, config AppConfig) context.Context {
	if ctx.Value(contextConfig) == nil {
		ctx = context.WithValue(ctx, contextConfig, config)
	}
	return ctx
}

func GetConfigFromContext(ctx context.Context) AppConfig {
	config, ok := ctx.Value(contextConfig).(AppConfig)

	if !ok {
		return AppConfig{}
	}

	return config
}
