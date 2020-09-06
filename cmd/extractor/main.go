package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/opencars/translit"

	"github.com/opencars/edrmvs/pkg/config"
	"github.com/opencars/edrmvs/pkg/hsc"
	"github.com/opencars/edrmvs/pkg/logger"
	"github.com/opencars/edrmvs/pkg/model"
	"github.com/opencars/edrmvs/pkg/store"
	"github.com/opencars/edrmvs/pkg/store/sqlstore"
)

func main() {
	var configPath, series string
	var from int64

	flag.StringVar(&configPath, "config", "./config/config.yaml", "Path to the configuration file")
	flag.StringVar(&series, "series", "", "Series of the registrations")
	flag.Int64Var(&from, "from", -1, "Use custom from")

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

	series = translit.ToLatin(series)
	reg, err := s.Registration().GetLast(series)
	if err != nil && err != store.ErrRecordNotFound {
		logger.Fatalf("store: last registration: %v", err)
	}

	var number int64
	if from != -1 {
		number = from
	} else if reg != nil {
		number, err = strconv.ParseInt(reg.NDoc, 10, 64)
		if err != nil {
			logger.Fatalf("convert registration number: %v", err)
		}
	}

	client := hsc.New(conf.HSC.BaseURL)
	for i := number; i < 1000000; i++ {
		code := fmt.Sprintf("%s%06d", series, i)

		l := logger.WithFields(logger.Fields{
			"code": code,
		})

		l.Debugf("sending request")

		regs, err := client.VehiclePassport(code)
		if err != nil {
			l.Errorf("request failed: %v", err)
			l.Debugf("sleep for 30s and then retry")
			time.Sleep(30 * time.Second)
			i--
			continue
		}

		if len(regs) > 1 {
			l.Errorf("too many registrations detected: %d", len(regs))
			continue
		}

		if len(regs) == 0 {
			l.Debugf("no registrations detected")
			continue
		}

		obj, err := model.FromHSC(regs[0])
		if err != nil {
			l.Errorf("convert: %s", err)
			l.Debugf("sleep for 30s and then retry")
			time.Sleep(30 * time.Second)
			continue
		}

		if err := s.Registration().Create(obj); err != nil {
			l.Fatalf("save registration: %v", err)
		}
	}
}
