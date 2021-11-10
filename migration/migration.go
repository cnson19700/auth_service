package migration

import (
	"context"
	"log"

	"github.com/pressly/goose"

	"github.com/auth_service/client/mysql"
	"github.com/auth_service/config"
)

func Up() {
	cfg := config.GetConfig()

	db := mysql.GetClient(context.Background())
	sqlDB, err := db.DB()

	if err != nil {
		log.Fatal("Fail to connect db", err)
	}

	if cfg.AutoMigrate {
		err := goose.SetDialect("mysql")
		if err != nil {
			panic(err)
		}

		err = goose.Run("up", sqlDB, "./migration")
		if err != nil {
			panic(err)
		}
	}
}
