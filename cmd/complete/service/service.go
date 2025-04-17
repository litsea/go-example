package service

import (
	"context"
	"time"

	"github.com/litsea/kit/graceful"
	log "github.com/litsea/log-slog"
	"github.com/spf13/viper"

	"github.com/litsea/go-example/config"
)

type demoService struct {
	stop chan struct{}
}

func New(v *viper.Viper) {
	demo1 := &demoService{
		stop: make(chan struct{}),
	}

	httpHandle(v)

	gracefulRun(v, log.Get(), demo1)
}

func gracefulRun(v *viper.Viper, l *log.Logger, srv graceful.Service) {
	g := graceful.New(
		graceful.WithService(srv),
		graceful.WithLogger(l),
		graceful.WithStopTimeout(v.GetDuration(config.KeyStopTimeout)),
	)

	err := g.Run(context.Background())
	if err != nil {
		l.Error("service.gracefulRun", "error", err)
	}

	// Wait for send event to Sentry when server start failed
	time.Sleep(3 * time.Second)
}

func (s *demoService) Name() string {
	return "demo"
}

func (s *demoService) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			log.Info("service.Start: context cancelled", "name", s.Name())
			return nil
		case <-s.stop:
			log.Info("service.Start: stop signal received", "name", s.Name())
			return nil
		default:
			log.Info("service.Start: loop", "name", s.Name())
			time.Sleep(2 * time.Second)
		}
	}
}

func (s *demoService) Stop(ctx context.Context) error {
	close(s.stop)

	interval := 100 * time.Millisecond
	timer := time.NewTimer(interval)
	defer timer.Stop()

	cleanCh := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 3)
		close(cleanCh)
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-cleanCh:
			log.Info("service.Stop: cleanup done", "name", s.Name())
			timer.Reset(interval)
			return nil
		}
	}
}
