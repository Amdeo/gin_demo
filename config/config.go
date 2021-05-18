package config

type ServerConfig struct {
	System   System   `json:"system" yaml:"system"`
	Database Database `json:"database" yaml:"database"`
	Zap      Zap      `json:"zap" yaml:"zap"`
	Redis    Redis    `json:"redis" yaml:"redis"`
}

type System struct {
	Port string `yaml:"port"` // 服务启动端口
}

type Database struct {
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	Dbname   string `json:"dbname" yaml:"db-name"`    // 数据库名
	Username string `json:"username" yaml:"username"` // 数据库用户名
	Password string `json:"password" yaml:"password"` // 数据库密码
}

type Zap struct {
	Level        string `json:"level" yaml:"level"`
	Filename     string `json:"filename" yaml:"filename"`
	MaxSize      int    `json:"maxSize" yaml:"maxSize"`
	MaxBackups   int    `json:"maxBackups" yaml:"maxBackups"`
	MaxAge       int    `json:"maxAge" yaml:"maxAge"`
	LogInConsole bool   `json:"logInConsole" yaml:"log-in-console"` // 输出控制台
}

type Redis struct {
	DB       int    `json:"db" yaml:"db"`
	Addr     string `json:"addr" yaml:"addr"`
	Password string `json:"password" yaml:"password"`
}
