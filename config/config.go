package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

// Module is config module
var Module = fx.Options(
	fx.Provide(
		New,
	),
)

type argvMeta struct {
	desc       string
	defaultVal string
}

// New returns a viper object.
// This object is used to read environment variables or command line arguments.
func New() (config *viper.Viper) {
	config = viper.New()

	confList := map[string]argvMeta{
		"env": {
			defaultVal: "development",
			desc:       "Environment",
		},
		"postgres_db": {
			defaultVal: "blogdb",
			desc:       "postgres db name",
		},
		"postgres_host": {
			defaultVal: "localhost",
			desc:       "postgres host",
		},
		"postgres_port": {
			defaultVal: "5432",
			desc:       "postgres port",
		},
		"postgres_user": {
			defaultVal: "postgres",
			desc:       "postgres username",
		},
		"postgres_password": {
			defaultVal: "",
			desc:       "postgres password",
		},
		"port": {
			defaultVal: "8765",
			desc:       "Port number of user API server",
		},
		"mode": {
			defaultVal: "server",
			desc:       "App mode eg. consumer, server, worker",
		},
		"log_level": {
			defaultVal: "debug",
			desc:       "Log level to be printed. List of log level by Priority - debug, info, warn, error, dpanic, panic, fatal",
		},
	}

	for key, meta := range confList {
		// automatic conversion of environment var key to `UPPER_CASE` will happen.
		config.BindEnv(key)

		// read command-line arguments
		pflag.String(key, meta.defaultVal, meta.desc)
	}

	pflag.Parse()
	config.BindPFlags(pflag.CommandLine)
	return
}
