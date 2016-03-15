package dphx

// Config loads and saves config.
type Config struct {
	Username   string `envconfig:"SSH_USER"`                            // SSH user
	Password   string `envconfig:"SSH_PASSWORD"`                        // SSH pass
	ServerAddr string `envconfig:"SSH_ADDR"`                            // SSH addr
	LocalAddr  string `envconfig:"SOCKS_ADDR",default:"127.0.0.1:1080"` // SOCKS listen addr
}
