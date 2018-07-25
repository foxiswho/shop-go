package conf

type redis struct {
	Server string `toml:"server"`
	Pwd    string `toml:"pwd"`
}
