package processor

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/opencars/translit"

	"github.com/opencars/edrmvs/pkg/domain"
	"github.com/opencars/edrmvs/pkg/domain/model"
	"github.com/opencars/edrmvs/pkg/logger"
)

const (
	MaxAmountPerSeries = 1000000
)

type Processor struct {
	store    domain.SystemRegistrationStore
	provider domain.RegistrationProvider

	delay, backoff time.Duration
}

func New(store domain.SystemRegistrationStore, p domain.RegistrationProvider, delay, backoff time.Duration) *Processor {
	return &Processor{
		store:    store,
		provider: p,
		delay:    delay,
		backoff:  backoff,
	}
}

func (p *Processor) Process(ctx context.Context, series string, from int64) error {
	latinSeries := translit.ToLatin(series)

	start, err := p.start(ctx, latinSeries, from)
	if err != nil {
		return err
	}

	delay := p.delay
	for i := start; i < MaxAmountPerSeries; i++ {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(delay):
			delay = p.delay
		}

		code := fmt.Sprintf("%s%06d", latinSeries, i)

		l := logger.WithFields(logger.Fields{
			"code": code,
		})

		l.Infof("sending request")

		regs, err := p.provider.FindByCode(ctx, code)
		if errors.Is(err, context.Canceled) {
			return nil
		}

		if err != nil {
			l.Errorf("request failed: %v", err)
			l.Debugf("sleep for %s and then retry", p.backoff)
			delay = p.backoff
			i--
			continue
		}

		if len(regs) > 1 {
			l.Errorf("too many registrations detected: %d", len(regs))
			continue
		}

		if len(regs) == 0 {
			l.Infof("no registrations detected")
			continue
		}

		for _, r := range regs {
			r.Number = translit.ToLatin(r.Number)
		}

		err = p.store.Create(ctx, &regs[0])
		if errors.Is(err, context.Canceled) {
			return nil
		}

		if err != nil {
			return err
		}

		l.Infof("successfully created")
	}

	return nil
}

func (p *Processor) start(ctx context.Context, series string, from int64) (int64, error) {
	if from != 0 {
		return from, nil
	}

	last, err := p.store.FindLastBySeries(ctx, series)
	if errors.Is(err, model.ErrNotFound) {
		return 0, nil
	}

	if err != nil {
		return 0, fmt.Errorf("last: %w", err)
	}

	return strconv.ParseInt(last.DocumentNumber, 10, 64)
}
