package db

// Config 数据库配置
type Config struct {
	DBType          string `toml:"DBType" yaml:"db_type"`     // 数据库驱动类型
	Host            string `toml:"Host" yaml:"host"`          // 主机地址
	Port            int    `toml:"Port" yaml:"port"`          // 端口号
	DBName          string `toml:"DBName" yaml:"db_name"`     // 数据库名
	UserName        string `toml:"UserName" yaml:"user_name"` // 用户名
	PassWord        string `toml:"PassWord" yaml:"pass_word"` // 密码
	CharSet         string `toml:"CharSet" yaml:"char_set"`   // 字符集
	MaxIdleConns    int    `toml:"MaxIdleConns" yaml:"max_idle_conns"`
	MaxOpenConns    int    `toml:"MaxOpenConns" yaml:"max_open_conns"`
	ConnMaxLifetime int    `toml:"ConnMaxLifetime" yaml:"conn_max_lifetime"`
	LogMode         bool   `toml:"LogMode" yaml:"log_mode"`
}
