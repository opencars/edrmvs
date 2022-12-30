package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/opencars/seedwork/logger"

	"github.com/opencars/edrmvs/pkg/config"
	"github.com/opencars/edrmvs/pkg/domain/processor"
	"github.com/opencars/edrmvs/pkg/hsc"
	"github.com/opencars/edrmvs/pkg/store/sqlstore"
)

func main() {
	configPath := flag.String("config", "./config/config.yaml", "Path to the configuration file")
	series := flag.String("series", "CXE", "Series of the registrations")
	from := flag.Int64("from", 0, "Use custom from")

	flag.Parse()

	conf, err := config.New(*configPath)
	if err != nil {
		logger.Fatalf("failed read config: %v", err)
	}

	logger.NewLogger(logger.LogLevel(conf.Log.Level), conf.Log.Mode == "dev")

	s, err := sqlstore.New(&conf.DB)
	if err != nil {
		logger.Fatalf("store: %v", err)
	}

	provider := hsc.NewProvider(hsc.New(&conf.HSC))
	p := processor.New(s, provider,
		conf.Extractor.Delay.Duration,
		conf.Extractor.BackOff.Duration,
	)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-c
		cancel()
	}()

	if err := p.Process(ctx, *series, *from); err != nil {
		logger.Fatalf("process: %s", err)
	}
}
