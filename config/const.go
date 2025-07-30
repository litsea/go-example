package config

const (
	KeyEnv = "env"

	// server

	KeyServerAddr  = "server.addr"
	KeyPprofToken  = "server.pprof-token" //nolint:gosec
	KeyStopTimeout = "server.stop-timeout"

	// profiler

	KeyProfilerServerAddress = "profiler.server-address"
	KeyProfilerAuthUsername  = "profiler.auth-username"
	KeyProfilerAuthPassword  = "profiler.auth-password" //nolint:gosec
	KeyProfilerDebug         = "profiler.debug"

	// metrics

	KeyMetricsEnable   = "metrics.enable"
	KeyMetricsUsername = "metrics.username"
	KeyMetricsPassword = "metrics.password"

	// log

	KeyLog = "log"
)
