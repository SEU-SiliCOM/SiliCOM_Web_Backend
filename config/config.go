package config

type Config struct {
	Port         string       `toml:"port"`
	Rm_rf        bool         `toml:"rm-rf"`
	BonusPath    string       `toml:"bonusPath"`
	Email        Email        `toml:"email"`
	Mysql        Mysql        `toml:"mysql"`
	Upload       Upload       `toml:"upload"`
	Jwt          Jwt          `toml:"jwt"`
	Code2Session Code2Session `toml:"code2Session"`
}

type Email struct {
	Address  string `toml:"address"`
	Name     string `toml:"name"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Password string `toml:"password"`
}

type Mysql struct {
	Username string `toml:"username"`
	Password string `toml:"password"`
	Ip       string `toml:"ip"`
	Port     string `toml:"port"`
	Database string `toml:"database"`
}

type Upload struct {
	Path      string `toml:"path"`
	AccessUrl string `toml:"accessUrl"`
}

type Jwt struct {
	Key string `toml:"key"`
}

type Code2Session struct {
	AppId     string `toml:"appId"`
	AppSecret string `toml:"appSecret"`
}
