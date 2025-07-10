package config

const (
	KeyEnv = "env"

	// server

	KeyPprofAddr   = "server.pprof-addr"
	KeyPprofToken  = "server.pprof-token" //nolint:gosec
	KeyStopTimeout = "server.stop-timeout"

	// profiler

	KeyProfilerServerAddress = "profiler.server-address"
	KeyProfilerAuthUsername  = "profiler.auth-username"
	KeyProfilerAuthPassword  = "profiler.auth-password" //nolint:gosec
	KeyProfilerDebug         = "profiler.debug"

	// log

	KeyLog = "log"
)
