package conf

type Configure struct {
	DBConfig struct {
		Host string
		Port int
		DbName string
		User string
		Password string
	}

	LogConf struct {
		LogType string //only raw

		Output struct {
			OutputType string
			Var string
		}
	}
}
