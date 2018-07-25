package conf

type opentracing struct {
	Disable     bool   `toml:"disable"`
	Type        string `toml:"type"`
	ServiceName string `toml:"service_name"`
	Address     string `toml:"address"`
}