package service

import (
	"net/http"
	"time"

	"github.com/litsea/kit/health"
	"github.com/litsea/kit/pprof"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"

	"github.com/litsea/go-example/config"
	"github.com/litsea/go-example/internal/metrics"
)

func httpHandle(v *viper.Viper) {
	mux := http.NewServeMux()

	pprof.Register(mux, func() string {
		return v.GetString(config.KeyPprofToken)
	})

	if config.V().GetBool(config.KeyMetricsEnable) {
		// Prometheus metrics
		mux.Handle("/metrics",
			metrics.AuthMiddleware(
				config.V().GetString(config.KeyMetricsUsername),
				config.V().GetString(config.KeyMetricsPassword),
				promhttp.Handler(),
			),
		)
		metrics.Init()
	}

	health.Register(mux, "/v1/health")

	s := &http.Server{
		Handler:           mux,
		Addr:              v.GetString(config.KeyServerAddr),
		ReadHeaderTimeout: 3 * time.Second,
	}

	go func() {
		_ = s.ListenAndServe()
	}()
}
