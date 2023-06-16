package config

var Config struct {
	Databases []Database `json:"databases"`
}

type Database struct {
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DB       string `json:"db"`
}

func GetDatabaseByDb(db string) *Database {
	for _, database := range Config.Databases {
		if database.DB == db {
			return &database
		}
	}
	return nil
}
