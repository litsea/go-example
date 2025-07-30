package config

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	log "github.com/litsea/log-slog"
	vp "github.com/litsea/viper-aws"
	"github.com/litsea/viper-aws/secrets"
	"github.com/spf13/viper"
)

var config *Config

type Config struct {
	v *viper.Viper
	p *secrets.Provider
}

func Init(cfgFile, cfgType string) {
	var (
		err error
		p   *secrets.Provider
		cfg *vp.Config
	)

	l := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
	v := viper.NewWithOptions(viper.WithLogger(l))
	v.SetEnvPrefix("GO_EXP")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()
	sid := os.Getenv("CONFIG_AWS_SECRET_ID")

	if sid == "" {
		cfg, err = vp.NewFile(v,
			vp.WithType(cfgType),
			vp.WithFile(cfgFile),
			vp.WithOnFileChange(onFileChange(v)),
			vp.WithSetDefaultFunc(setDefault),
		)
	} else {
		cfg, err = vp.NewSecrets(v, sid,
			[]vp.Option{vp.WithType(cfgType)},
			[]secrets.Option{
				secrets.WithLogger(l),
				secrets.WithOnChangeFunc(onSecretsChange(v, cfgType)),
			},
		)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	config = &Config{v: cfg.V(), p: p}

	err = InitLogger(v)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	*l = *log.GetSlog()
}

func V() *viper.Viper {
	return config.v
}
