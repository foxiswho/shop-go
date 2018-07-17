package conf

type server struct {
	Graceful bool   `toml:"graceful"`
	Addr     string `toml:"addr"`

	DomainApi    string `toml:"domain_api"`
	DomainAdmin  string `toml:"domain_admin"`
	DomainWeb    string `toml:"domain_web"`
	DomainSocket string `toml:"domain_socket"`
}
