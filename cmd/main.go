package main

import (
	"log"
	"net"
	"time"

	"github.com/Auth-Service/client/mysql"
	"github.com/Auth-Service/config"
	"github.com/Auth-Service/usecase"

	serviceHttp "github.com/Auth-Service/delivery/http"
	"github.com/Auth-Service/migration"
	"github.com/Auth-Service/repository"
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
