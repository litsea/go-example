package metrics

import (
	"net/http"
	"sync"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/litsea/go-example/version"
)

var (
	prefix = "goexample_"
	once   sync.Once

	versionInfo = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: prefix + "app_version_info",
			Help: "Application version info as labels",
		},
		[]string{"version", "git_rev", "git_branch", "build_date"},
	)
)

func Init() {
	once.Do(func() {
		prometheus.MustRegister(
			versionInfo,
		)

		versionInfo.WithLabelValues(
			version.Version, version.GitRev, version.GitBranch, version.BuildDate,
		).Set(1)
	})
}

func AuthMiddleware(username, password string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()
		if !ok || u != username || p != password {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
