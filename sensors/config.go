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

//Mqtt struct for mqtt settings
type Mqtt struct {
	User        string
	Password    string
	Host        string
	Port        string
	Scheme      string
	ClientID    string
	TopicInput  string
	TopicOutput string
}

//Config struct with all different configuration entries
type Config struct {
	DB     Database
	Server Server
	Mqtt   Mqtt
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
		Mqtt: Mqtt{
			User:        cfg.Mqtt.User,
			Password:    cfg.Mqtt.Password,
			Host:        cfg.Mqtt.Host,
			Port:        cfg.Mqtt.Port,
			Scheme:      cfg.Mqtt.Scheme,
			ClientID:    cfg.Mqtt.ClientID,
			TopicInput:  cfg.Mqtt.TopicInput,
			TopicOutput: cfg.Mqtt.TopicOutput,
		},
		Secret: cfg.AppSecret,
	}
}
