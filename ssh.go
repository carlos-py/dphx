package dphx

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
	"golang.org/x/net/context"
)

var sshClient *ssh.Client

func createSSHClient() (*ssh.Client, error) {
	sshConfig, err := appConfig.SSH.ClientConfig()

	if err != nil {
		return nil, err
	}

	log.Printf("Connecting to SSH at %s", appConfig.SSH.ServerAddr)

	client, err := ssh.Dial("tcp", appConfig.SSH.ServerAddr, sshConfig)
	if err != nil {
		return nil, fmt.Errorf("Failed to dial to SSH server: %s", err.Error())
	}

	return client, nil
}

func ensureSSHClient() error {
	if sshClient == nil {
		newClient, err := createSSHClient()
		if err != nil {
			return err
		}
		sshClient = newClient
	}
	return nil
}

func addressForDial(ctx context.Context, addr string) (string, error) {
	// Parse the address into host and numeric port.
	host, portString, err := net.SplitHostPort(addr)
	if err != nil {
		return "", err
	}

	// Replace the address with the original address before local
	// DNS resolution.
	if newHost := fetchOriginalAddress(ctx); newHost != "" {
		host = newHost
	}

	// Join new host and port back.
	addr = net.JoinHostPort(host, portString)

	return addr, nil
}

// SSHDial does dial via SSH connection to a remote server.
func SSHDial(ctx context.Context, network, addr string) (net.Conn, error) {
	if err := ensureSSHClient(); err != nil {
		return nil, err
	}

	addr, err := addressForDial(ctx, addr)
	if err != nil {
		return nil, err
	}

	return sshClient.Dial(network, addr)
}
