package config

import (
	"fmt"
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

//DatabaseConfig has database config info
type DatabaseConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"dbname"`
}

//ServerConfig has server configuration options
type ServerConfig struct {
	Port string `yaml:"port"`
}

//MqttConfig has Mqtt configuration options
type MqttConfig struct {
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	Scheme      string `yaml:"scheme,omitempty"`
	ClientID    string `yaml:"client_id"`
	TopicInput  string `yaml:"topic_input"`
	TopicOutput string `yaml:"topic_output"`
}

//YamlConfig maps the configuration in the yaml file
//into a struct
type YamlConfig struct {
	Database  DatabaseConfig `yaml:"database"`
	Server    ServerConfig   `yaml:"server"`
	Mqtt      MqttConfig     `yaml:"mqtt"`
	AppSecret string         `yaml::app_secret_key`
}

//ReadYaml will load a yaml file from the given path and
//return a YamlConfig struct
func ReadYaml(path string) *YamlConfig {
	if path == "" {
		path = defaultYamlConfigPath()
	}

	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	defer func() { _ = f.Close() }()

	var cfg YamlConfig
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Printf("error decoding config yaml file into config: %s\n", err)
		os.Exit(2)
	}

	return &cfg
}

func defaultYamlConfigPath() string {
	ex, err := os.Executable()
	if err != nil {
		log.Printf("error encountered reading path: %s\n", err)
		os.Exit(2)
	}

	filename := "config.yaml"
	dir := path.Dir(path.Dir(ex))
	dir = path.Join(dir, filename)
	return dir
}
