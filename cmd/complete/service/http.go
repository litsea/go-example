package service

import (
	"net/http"
	"time"

	"github.com/litsea/kit/health"
	"github.com/litsea/kit/pprof"
	"github.com/spf13/viper"

	"github.com/litsea/go-example/config"
)

func httpHandle(v *viper.Viper) {
	mux := http.NewServeMux()

	pprof.Register(mux, func() string {
		return v.GetString(config.KeyPprofToken)
	})
	health.Register(mux, "/v1/health")

	s := &http.Server{
		Handler:           mux,
		Addr:              v.GetString(config.KeyPprofAddr),
		ReadHeaderTimeout: 3 * time.Second,
	}

	go func() {
		_ = s.ListenAndServe()
	}()
}
