package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/parsa-poorsistani/http-monitoring-system/pkg/config"
	"github.com/parsa-poorsistani/http-monitoring-system/pkg/database"
	"github.com/parsa-poorsistani/http-monitoring-system/pkg/healthcheck"
	"github.com/parsa-poorsistani/http-monitoring-system/pkg/server"

	"net/http"
)
var log = logrus.New()

func init() {
  file, err := os.OpenFile("/root/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

  if err == nil {
    log.Out = file
  } else {
    logrus.Info("Failed to log to file, using default stderr")
  }
  log.SetLevel(logrus.InfoLevel)

  log.SetFormatter(&logrus.JSONFormatter{})
}
func main() {
  log.Info("App starting ...")
  fmt.Print("APP Starting ...")

  cfg := config.LoadConfig() 
  //if err != nil {
    //fmt.Printf("config load fucked ...:%v \n", err)
    //log.WithError(err).Fatal("Failed to load config file")
  //}

 // dbUser, err := os.ReadFile("/etc/secret-volume/username")
  //if err != nil {
    //  log.Fatal("Failed to read database user secret", err)
  //}
  //dbPassword, err := os.ReadFile("/etc/secret-volume/password")
  //if err != nil {
    //  log.Fatal("Failed to read database password secret", err)
  //}

  //cfg.Database.User = string(dbUser)
  //cfg.Database.Password = string(dbPassword)

  db, err := database.NewDatabase(cfg)
  if err != nil {
    fmt.Printf("DB FUCKED ...: %v \n", err)
    log.WithError(err).Fatal("Failed to set up the DB")
  }
  db.InitDb()
  srv := server.NewServer(db,cfg,log)

  httpAddr := ":" + cfg.Server.Port
  log.WithField("port", cfg.Server.Port).Info("Starting HTTP Server")
  if err := http.ListenAndServe(httpAddr, srv.Router()); err != nil {
    fmt.Print("server fucked ...")
    log.WithError(err).Fatal("Failed to start HTTP Srver")
  }

  healthchecker := healthcheck.NewHealthChecker(db, cfg, log)
  go healthchecker.Start()
}
