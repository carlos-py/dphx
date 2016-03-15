package dphx

// Config loads and saves config.
type Config struct {
	Username   string // SSH user
	Password   string // SSH pass
	ServerAddr string // SSH addr
	LocalAddr  string // SOCKS listen addr
}
