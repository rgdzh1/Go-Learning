package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName  string         `json:"app_name"`
	AppMode  string         `json:"app_mode"`
	AppHost  string         `json:"app_host"`
	AppPort  string         `json:"app_port"`
	Sms      SmsConfig      `json:"sms"`
	DataBase DataBaseConfig `json:"database"`
}
type SmsConfig struct {
	SignName     string `json:"sign_name"`
	TemplateCode string `json:"template_code"`
	AppKey       string `json:"app_key"`
	AppSecret    string `json:"app_secret"`
	RegionId     string `json:"region_id"`
}
type DataBaseConfig struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DBName   string `json:"db_name"`
	Charset  string `json:"charset"`
	ShowSql  bool   `json:"show_sql"`
}

// _号,禁止访问
var _cfg = new(Config)

/**
返回Config对象
*/
func GetConfig() *Config {
	return _cfg
}
func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(_cfg)
	if err != nil {
		return nil, err
	}
	return _cfg, nil

}