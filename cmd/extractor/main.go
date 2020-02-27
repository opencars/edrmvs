package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/lib/pq"

	"github.com/opencars/edrmvs/pkg/config"
	"github.com/opencars/edrmvs/pkg/hsc"
	"github.com/opencars/edrmvs/pkg/model"
	"github.com/opencars/edrmvs/pkg/store"
	"github.com/opencars/edrmvs/pkg/store/sqlstore"
	"github.com/opencars/translit"
)

func main() {
	var path, series string
	var from int64

	flag.StringVar(&path, "config", "./config/config.toml", "Path to the configuration file")
	flag.StringVar(&series, "series", "", "Series of the registrations")
	flag.Int64Var(&from, "from", -1, "Use custom from")

	flag.Parse()

	settings, err := config.New(path)
	if err != nil {
		log.Fatal(err)
	}

	s, err := sqlstore.New(settings)
	if err != nil {
		log.Fatal(err)
	}

	series = translit.ToLatin(series)
	reg, err := s.Registration().GetLast(series)
	if err != nil && err != store.ErrRecordNotFound {
		log.Fatal(err)
	}

	var number int64 = 0
	if from != -1 {
		number = from
	} else if reg != nil {
		number, err = strconv.ParseInt(reg.NDoc, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
	}

	client := hsc.New(settings.HSC.BaseURL)
	for i := number; i < 1000000; i++ {
		code := fmt.Sprintf("%s%06d", series, i)
		fmt.Println(code)

		regs, err := client.VehiclePassport(code)
		if err != nil {
			fmt.Printf("[warn] Error on code=%s: %s\n", code, err)
			time.Sleep(30 * time.Second)
			i--
			continue
		}

		if len(regs) > 1 {
			fmt.Printf("[warn] Too many regs on code=%s\n", code)
			continue
		}

		if len(regs) == 0 {
			continue
		}

		obj, err := model.FromHSC(regs[0])
		if err != nil {
			fmt.Printf("[warn] FromHSC failed on code=%s: %s\n", code, err)
			time.Sleep(30 * time.Second)
			continue
		}

		if err := s.Registration().Create(obj); err != nil {
			log.Fatal(err)
		}
	}
}
