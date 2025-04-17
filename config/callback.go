package config

import (
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/fsnotify/fsnotify"
	"github.com/litsea/log-slog"
	"github.com/spf13/viper"
)

func setDefault(v *viper.Viper) {
	v.SetDefault(KeyPprofAddr, "0.0.0.0:8080")
	v.SetDefault(KeyStopTimeout, 30*time.Second)
}

func onFileChange(_ *viper.Viper) func(evt fsnotify.Event) {
	return func(evt fsnotify.Event) {
		log.Warn("config.onFileChange", "name", evt.Name, "op", evt.Op.String())

		// Remove comment for testing
		// panic("config.onFileChange: panic test")
	}
}

func onSecretsChange(_ *viper.Viper, cfgType string) func(out *secretsmanager.GetSecretValueOutput) {
	return func(out *secretsmanager.GetSecretValueOutput) {
		// Check syntax
		vv := viper.New()
		vv.SetConfigType(cfgType)
		err := vv.ReadConfig(strings.NewReader(*out.SecretString))
		if err != nil {
			log.Error("config.onSecretsChange: invalid value", "version", *out.VersionId,
				"createdDate", out.CreatedDate, "err", err)
			return
		}

		log.Warn("config.onSecretsChange: updated", "version", *out.VersionId,
			"createdDate", out.CreatedDate)

		// Remove comment for testing
		// panic("config.onSecretsChange: panic test")
	}
}

func QuitWatch() {
	if config != nil && config.p != nil {
		config.p.QuitWatch()
	}
}
