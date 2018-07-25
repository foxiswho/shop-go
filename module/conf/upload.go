package conf

type Upload struct {
	Type        string `toml:"type"`          //上传方式 local:本地 QiNiu:七牛云存储
	Ext         string `toml:"ext"`           //允许上传后缀
	RootPath    string `toml:"root_path"`     //上传文件目录
	RootPathTmp string `toml:"root_path_tmp"` //临时文件目录
	Size        int    `toml:"size"`          //最大上传文件大小 5*1024*1024
	LocalSaveIs bool   `toml:"local_save_is"` //是否本地保存
	Http        string `toml:"http"`          //域名
}