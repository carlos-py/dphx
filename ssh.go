package dphx

import (
	"fmt"
	"net"

	"golang.org/x/crypto/ssh"
	"golang.org/x/net/context"
)

var sshClient *ssh.Client

func createSSHClient() (*ssh.Client, error) {
	// Only password auth is supported so far.
	sshConfig := &ssh.ClientConfig{
		User: appConfig.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(appConfig.Password),
		},
	}
	client, err := ssh.Dial("tcp", appConfig.ServerAddr, sshConfig)
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

// SSHDial does dial via SSH connection to a remote server.
func SSHDial(ctx context.Context, network, addr string) (net.Conn, error) {
	if err := ensureSSHClient(); err != nil {
		return nil, err
	}

	// Parse the address into host and numeric port.
	host, portString, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}

	// Replace the address with the original address before local
	// DNS resolution.
	originalAddress := ctx.Value("originalAddress")
	if originalAddress != nil {
		originalAddressTyped, ok := originalAddress.(OriginalAddress)
		if ok {
			host = originalAddressTyped.String()
		}
	}

	addr = net.JoinHostPort(host, portString)

	return sshClient.Dial(network, addr)
}
