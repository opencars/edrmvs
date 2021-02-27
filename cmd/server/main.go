package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/opencars/edrmvs/pkg/api/http"
	"github.com/opencars/edrmvs/pkg/config"
	"github.com/opencars/edrmvs/pkg/domain/registration"
	"github.com/opencars/edrmvs/pkg/hsc"
	"github.com/opencars/edrmvs/pkg/logger"
	"github.com/opencars/edrmvs/pkg/store/sqlstore"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "./config/config.yaml", "Path to the configuration file")

	flag.Parse()
	conf, err := config.New(configPath)
	if err != nil {
		logger.Fatalf("failed read config: %v", err)
	}

	logger.NewLogger(logger.LogLevel(conf.Log.Level), conf.Log.Mode == "dev")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-c
		cancel()
	}()

	s, err := sqlstore.New(&conf.DB)
	if err != nil {
		logger.Fatalf("store: %v", err)
	}

	p := hsc.NewProvider(hsc.New(&conf.HSC))
	r := registration.NewService(s, p)

	addr := ":8080"
	logger.Infof("Listening on %s...", addr)
	if err := http.Start(ctx, addr, &conf.Server, r); err != nil {
		logger.Fatalf("http server failed: %v", err)
	}
}
