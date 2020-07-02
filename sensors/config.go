package sensors

import (
	"fmt"
	"sensors/config"
)

//Database struct for db settings
type Database struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func (d Database) String(sslmode string) string {
	return fmt.Sprintf(""+
		"host=%s "+
		"port=%s "+
		"user=%s "+
		"dbname=%s "+
		"password=%s "+
		"sslmode=%v "+
		"", d.Host, d.Port, d.User, d.DBName, d.Password, sslmode)
}

//Server struct for server settings
type Server struct {
	Port string
}

//Config struct with all different configuration entries
type Config struct {
	DB Database

	Server Server

	Secret string
}

//GetConfig returns a configuration struct
func GetConfig(cfg config.YamlConfig) Config {
	return Config{
		DB: Database{
			User:     cfg.Database.User,
			Password: cfg.Database.Password,
			Host:     cfg.Database.Host,
			Port:     cfg.Database.Port,
			DBName:   cfg.Database.DBName,
		},

		Server: Server{
			Port: cfg.Server.Port,
		},

		Secret: cfg.AppSecret,
	}
}
