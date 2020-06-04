package conf

type DBConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	Running  bool
}

// DBList : list of all mysql master services
var DBList = []DBConfig{
	{
		Host:     "127.0.0.1",
		Port:     13306,
		Username: "root",
		Password: "123456",
		Database: "event",
		Running:  true,
	},
}

// DBMain : Main db config
var DBMain = DBList[0]
