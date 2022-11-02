package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/YuLaFreitas/ConversoDB/configs"
	"github.com/sirupsen/logrus"

	conversordb "github.com/YuLaFreitas/ConversoDB/framework/gateways/repository/conversor_db"
)

var (
	log = logrus.WithField("package", "main")
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
}

func main() {
	ctx := setupContext()

	log := log.WithContext(ctx).WithField("main", ctx)

	conversordb := setupRepostory(ctx)

	fmt.Print(conversordb)

	log.Debug("applications started")

}

func setupContext() context.Context {
	ctx := context.Background()

	config := configs.AppConfig{}

	return configs.NewContextWithConfig(ctx, config)
}

func setupRepostory(ctx context.Context) conversordb.ConversoDB {
	config := configs.GetConfigFromContext(ctx)

	conversordb_repository, err := conversordb.NewMysqlRepository(config.Mysql)

	if err != nil {
		logrus.WithError(err)
	}

	return *conversordb_repository

}
