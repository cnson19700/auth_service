package main

import (
	"log"
	"net"
	"time"

	"github.com/auth_service/client/mysql"
	"github.com/auth_service/config"
	"github.com/auth_service/usecase"

	serviceHttp "github.com/auth_service/delivery/http"
	"github.com/auth_service/migration"
	"github.com/auth_service/repository"
)

func main() {
	cfg := config.GetConfig()

	// setup locale
	{
		loc, _ := time.LoadLocation(cfg.TimeZone)
		time.Local = loc
	}

	mysql.Init()
	migration.Up()

	repo := repository.New(mysql.GetClient)
	ucase := usecase.New(repo)

	executeServer(repo, ucase)

}

func executeServer(repo *repository.Repository, ucase *usecase.UseCase) {
	cfg := config.GetConfig()

	l, err := net.Listen("tcp", ":"+cfg.Port)

	if err != nil {
		log.Fatal(err)
	}

	errs := make(chan error)

	// http
	h := serviceHttp.NewHTTPHandler(repo, ucase)

	go func() {
		h.Listener = l

		log.Printf("Server is running on http://localhost:%s", cfg.Port)
		errs <- h.Start("")
	}()

	log.Println("exit", <-errs)
}
