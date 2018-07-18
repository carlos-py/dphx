package sshKraken // import "github.com/carlos-py/sshKraken"

import (
	"fmt"
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

func getKeyFile(filename string) (key ssh.Signer) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	key, err = ssh.ParsePrivateKey(buf)
	if err != nil {
		return
	}
	return
}

func Challenge(user, instruction string, questions []string, echos []bool) (answers []string, err error) {
	var suitable_answers = []string{}
	var pwIdx = 0
	answers = make([]string, len(questions))
	for n, q := range questions {
		fmt.Printf("Got question: %s\n", q)
		answers[n] = suitable_answers[pwIdx]
	}
	pwIdx++

	return answers, nil
}

// ClientConfig creates ssh.ClientConfig object from current configuration.
func (s SshServer) ClientConfig() (*ssh.ClientConfig, error) {

	sshConfig := &ssh.ClientConfig{
		User: s.Username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(getKeyFile(s.Key)),
			ssh.KeyboardInteractive(Challenge),
		},
	}

	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	return sshConfig, nil
}
