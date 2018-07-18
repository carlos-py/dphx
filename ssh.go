package sshKraken // import "github.com/carlos-py/sshKraken"

import (
	"fmt"
	"log"
	"net"
	"strings"

	"golang.org/x/crypto/ssh"
	"golang.org/x/net/context"
)

// SSHClients maps an SSH client connection to an index
type SSHClients struct {
	client map[int]*CountingSSHClient
}

var sshClientMap = SSHClients{client: make(map[int]*CountingSSHClient)}

// connectSSHServer attempts to establish an SSH connection to the server
// configured at index position provided
func (c *SSHClients) connectSSHServer(sshClient *CountingSSHClient, index int) error {
	if sshClient == nil {
		sshConfig, _ := appConfig.SSH[index].ClientConfig()
		log.Printf("[+] SSH Tunnel #%v Connecting to SSH at %s", index, appConfig.SSH[index].Host)
		client, err := ssh.Dial("tcp", appConfig.SSH[index].Host, sshConfig)
		if err != nil {
			return fmt.Errorf("Failed to dial to SSH server: %s", err.Error())
		}
		sshClientMap.client[index] = NewCountingSSHClient(client)
		return nil
	}
	return nil
}

// EnsureSSHClients loops through each SSH server configured and
// determines whether the SSH connection is active
func EnsureSSHClients() {
	for configIndex := range appConfig.SSH {

		if sshClientMap.client[configIndex] == nil {
			var newSSHClient *CountingSSHClient
			sshClientMap.connectSSHServer(newSSHClient, configIndex)
		}
	}
}

// SSHDial does dial via SSH connection to a remote server.
// The URL is checked against each 'match' string in SSH configuration
func SSHDial(ctx context.Context, network, addr string) (net.Conn, error) {
	for configIndex := range appConfig.SSH {
		switch {
		case sshClientMap.client[configIndex] != nil && strings.Contains(addr, appConfig.SSH[configIndex].URLMatch):

			log.Println("[+] Matched route: ", appConfig.SSH[configIndex].URLMatch, " via ", appConfig.SSH[configIndex].Host)

			connnectionDial, err := sshClientMap.client[configIndex].Dial(network, addr)

			if err != nil {
				sshClientMap.client[configIndex] = nil
				EnsureSSHClients()
			}
			return connnectionDial, err
		}
	}

	// Default route: If no URL match, route to the first SSH server configured
	connectionDial, err := sshClientMap.client[0].Dial(network, addr)
	if err != nil {
		sshClientMap.client[0] = nil
		EnsureSSHClients()
	}
	return connectionDial, err
}
