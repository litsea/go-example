package service

import (
	"context"
	"time"

	"github.com/litsea/kit/graceful"
	"github.com/litsea/kit/profiler"
	log "github.com/litsea/log-slog"
	"github.com/spf13/viper"

	"github.com/litsea/go-example/config"
)

type demoService struct {
	stop chan struct{}
}

func New(v *viper.Viper) {
	// Enable the profiler only when the server address provided
	if v.GetString(config.KeyProfilerServerAddress) != "" {
		_, err := profiler.Start(
			// appName.serviceName.env
			"go-example.complete."+v.GetString(config.KeyEnv),
			v.GetString(config.KeyProfilerServerAddress),
			profiler.WithAuth(
				v.GetString(config.KeyProfilerAuthUsername),
				v.GetString(config.KeyProfilerAuthPassword),
			),
			profiler.WithDebug(v.GetBool(config.KeyProfilerDebug)),
		)
		if err != nil {
			log.Error("service.New: failed to start profiler", "err", err)
		}
	}

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
			log.Info("service.Start: context cancelled", "service", s.Name())
			return nil
		case <-s.stop:
			log.Info("service.Start: stop signal received", "service", s.Name())
			return nil
		default:
			log.Info("service.Start: loop", "service", s.Name())
			time.Sleep(2 * time.Second)
		}
	}
}

func (s *demoService) Stop(ctx context.Context) error {
	close(s.stop)

	log.Info("service.Stop: service shutdown, cleaning in progress", "service", s.Name())
	time.Sleep(time.Second * 5)
	log.Info("service.Stop: service cleanup completed", "service", s.Name())

	return nil
}
