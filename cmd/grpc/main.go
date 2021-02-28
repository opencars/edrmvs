package main

import (
	"flag"
	"github.com/opencars/edrmvs/pkg/api/grpc"
	"github.com/opencars/edrmvs/pkg/config"
	"github.com/opencars/edrmvs/pkg/domain/registration"
	"github.com/opencars/edrmvs/pkg/hsc"
	"github.com/opencars/edrmvs/pkg/logger"
	"github.com/opencars/edrmvs/pkg/store/sqlstore"
	"os"
	"os/signal"
	"syscall"
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

	s, err := sqlstore.New(&conf.DB)
	if err != nil {
		logger.Fatalf("store: %v", err)
	}

	p := hsc.NewProvider(hsc.New(&conf.HSC))
	r := registration.NewService(s, p)

	addr := ":3000"
	api := grpc.New(addr, r)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-c
		if err := api.Close(); err != nil {
			logger.Fatalf("close: %v", err)
		}
	}()

	logger.Infof("Listening on %s...", addr)
	if err := api.Run(); err != nil {
		logger.Fatalf("grpc: %v", err)
	}
}
