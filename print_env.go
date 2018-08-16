package sshKraken // import "github.com/cvasq/sshKraken"

import (
	"fmt"
)

func printAppConfig(cfg AppConfig) {

	fmt.Printf("Starting sshKraken... \n\n")
	countServers := 0
	for _, host := range cfg.SSH {
		countServers++
		fmt.Printf("SSH tunnel #%v configuration \n", countServers)
		fmt.Printf("\t Hostname: %v \n", host.Host)
		fmt.Printf("\t Username: %v \n", host.Username)
		fmt.Printf("\t SSH Key: %v \n", host.Key)
		fmt.Printf("\t URL Match: %v \n\n", host.URLMatch)

	}

}
