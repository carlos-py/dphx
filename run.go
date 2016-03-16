package dphx

import (
	"github.com/kelseyhightower/envconfig"
)

// Run initializes config and starts SOCKS server.
func Run() {
	SetEnv()
	PrintEnv()
	ListenAndServe()
}

// SetEnv loads current env config.
func SetEnv() {
	envconfig.MustProcess("DPHX", &appConfig)
	envconfig.MustProcess("DPHX_SSH", &appConfig.SSH)
	SetSSHConfigDefaults(&appConfig.SSH)
}
