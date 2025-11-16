package config

type Config struct {
	BotToken        string
	DBType          string `yaml:"db_type"`
	DBPath          string `yaml:"db_path"`
	NCO_DBTableName string `yaml:"db_table_name"`
	Port            string `yaml:"port"`
	Address         string `yaml:"address"`
	LogFile         string `yaml:"log_file"`
}

func DefaultConfig() Config {
	return Config{
		NCO_DBTableName: "nco",
		DBPath:          "./data.db",
		LogFile:         "logs.log",
		Port:            "8080",
		Address:         "localhost",
	}
}
