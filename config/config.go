package config

type Config struct {
	BotToken        string
	DBType          string `yaml:"db_type"`
	DBPath          string `yaml:"db_path"`
	NCO_DBTableName string `yaml:"db_table_name"`
	Port            int    `yaml:"port"`
	LogFile         string `yaml:"log_file"`
}

func DefaultConfig() Config {
	return Config{
		NCO_DBTableName: "nco",
		DBPath:          "./data.db",
		LogFile:         "logs.log",
	}
}
