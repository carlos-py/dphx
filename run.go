package sshKraken // import "github.com/cvasq/sshKraken"

import (
	"log"
)

// Run initializes config and starts SOCKS server.
func Run() {
	config := LoadConfiguration("./config.json")

	appConfig.LocalAddr = "127.0.0.1:1080"
	appConfig.SSH = config

	printAppConfig(appConfig)

	go EnsureSSHClients()

	log.Printf("Local SOCKS5 server is starting at %s", appConfig.LocalAddr)
	if err := ListenAndServe("tcp", appConfig.LocalAddr); err != nil {
		log.Fatalln(err.Error())
	}
}
