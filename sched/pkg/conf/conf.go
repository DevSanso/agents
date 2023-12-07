package conf



const (
	LogConfDisplayType = "Display"
	LogConfFileType = "File"
)

type Configure struct {
	DBConfig struct {
		Host string
		Port int
		DbName string
		User string
		Password string
	}

	OdbcDriver string
	ScriptOption string

	LogConf struct {
		LogType string 
		LogVar string

		Output struct {
			OutputType string
			Var string
		}
	}
}
