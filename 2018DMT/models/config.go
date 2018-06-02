package models

//配置文件的模型
type Config struct {
	//网站根目录
	Wwwroot string `json:wwwroot`

	//端口
	Port int `json:port`
}
