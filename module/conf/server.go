package conf

type server struct {
	Graceful bool   `toml:"graceful"`
	Addr     string `toml:"addr"`

	DomainApi     string `toml:"domain_api"`
	DomainAdmin   string `toml:"domain_admin"`
	DomainWww     string `toml:"domain_www"`
	DomainSocket  string `toml:"domain_socket"`
	DomainExample string `toml:"domain_example"`
}
