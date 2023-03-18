package main

import (
	"context"
	"flag"
	"log"
	"main/crypter"
	"main/crypter/MD5Crypter"
	"main/linksService"
	"main/repository"
	"main/repository/postgre"
	"main/repository/storage"
	"main/server"
	"main/settings"
	"os"
	"os/signal"
	"syscall"
)

// TODO: move to config file
const (
	grpcPort = "50070"
	httpPort = "8080"
)

type app struct {
	httpServer server.HttpServer
	grpcServer server.GrpcServer
}

func (a app) start() {
	a.httpServer.Start()
	a.grpcServer.Start()
}
func (a app) shutdown() error {
	a.grpcServer.Stop()
	return a.httpServer.Stop()
}
func newApp(storage repository.IRepository, crypter crypter.ICrypter) (app, error) {
	linksServ := linksService.NewLinksService(storage, crypter)
	gs, err := server.NewGrpcServer(linksServ, grpcPort)
	if err != nil {
		return app{}, err
	}

	return app{
		httpServer: server.NewHttpServer(linksServ, httpPort),
		grpcServer: gs,
	}, nil
}
func run(ctx context.Context) error {
	repType := flag.Bool("d", false, "repository type")
	flag.Parse()
	var rep repository.IRepository
	if *repType {
		settings, err := settings.ParseSettingsFile("./db_credentials.json")
		if err != nil {
			return err
		}
		rep, err = postgre.Init(settings.HostName, settings.Port, settings.DBName, settings.Username, settings.Password)
		if err != nil {
			return err
		}
	} else {
		rep = storage.Init()
	}
	defer rep.Close()

	crypter := MD5Crypter.MD5Crypter{}
	app, err := newApp(rep, crypter)
	if err != nil {
		return err
	}
	app.start()
	defer app.shutdown()

	select {
	case httpErr := <-app.httpServer.Error():
		return httpErr
	case grpcErr := <-app.grpcServer.Error():
		return grpcErr
	case <-ctx.Done():
		return nil
	}
}
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), []os.Signal{os.Interrupt, syscall.SIGTERM}...)
	defer stop()
	if err := run(ctx); err != nil {
		log.Println(err)
		stop()
	}
}
