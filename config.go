package sshKraken // import "github.com/carlos-py/sshKraken"
import (
	"encoding/json"
	"fmt"
	"os"
)

// AppConfig loads and saves config.
type AppConfig struct {
	SSH       []SshServer
	LocalAddr string
}

type SshServer struct {
	Host      string `json:"host"`
	Username  string `json:"username"`
	Key       string `json:"key"`
	URLMatch  string `json:"urlMatch"`
	sshClient *CountingSSHClient
}

func LoadConfiguration(file string) []SshServer {
	var config []SshServer
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	appConfig.SSH = config
	return config
}
