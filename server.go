package dphx

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/MOZGIII/go-socks5"
	"golang.org/x/crypto/ssh"
	"golang.org/x/net/context"
)

var (
	username   = os.Getenv("DPHX_SSH_USER")
	password   = ssh.Password(os.Getenv("DPHX_SSH_PASSWORD"))
	serverAddr = os.Getenv("DPHX_SSH_ADDR")
	localAddr  = os.Getenv("DPHX_SOCKS_ADDR")
)

var sshClient *ssh.Client

func createSSHClient() (*ssh.Client, error) {
	// Only password auth is supported so far.
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			password,
		},
	}
	client, err := ssh.Dial("tcp", serverAddr, config)
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

// RemoteResolver resolved hostnames with remote server's sshd.
type RemoteResolver struct{}

// OriginalAddress preserves raw address for delayed DNS resolution.
type OriginalAddress string

func (o OriginalAddress) String() string {
	return string(o)
}

// Resolve implements resolution.
func (r RemoteResolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
	if err := ensureSSHClient(); err != nil {
		return ctx, nil, err
	}

	ctx = context.WithValue(ctx, "originalAddress", OriginalAddress(name))

	return ctx, net.IP{}, nil
}

func ListenAndServe() {
	// Create a SOCKS5 server
	conf := &socks5.Config{
		Dial:     SSHDial,
		Resolver: RemoteResolver{},
	}
	server, err := socks5.New(conf)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("SOCKS5 server is starting at %s", localAddr)
	if err := server.ListenAndServe("tcp", localAddr); err != nil {
		log.Fatalf(err.Error())
	}
}
