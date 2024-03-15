package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type execConfig struct {
	doMiguration bool `yaml:"do_miguration"`
}

func (ec execConfig) DoMiguration() bool {
	return ec.doMiguration
}

type sqlConfig struct {
	user string `yaml:"user"`
	pass string `yaml:"pass"`
	host string `yaml:"host"`
	name string `yaml:"name"`
	port uint   `yaml:"port"`
}

func (sc sqlConfig) User() string {
	return sc.user
}
func (sc sqlConfig) Pass() string {
	return sc.pass
}
func (sc sqlConfig) Host() string {
	return sc.host
}
func (sc sqlConfig) Name() string {
	return sc.name
}
func (sc sqlConfig) Port() uint {
	return sc.port
}

type appConfig struct {
	execConfig `yaml:"exec_config"`
	sqlConfig  `yaml:"sql_config"`
}

func (ac appConfig) ExecConfig() ExecConfig {
	return ac.execConfig
}

func (ac appConfig) SQLConfig() SQLConfig {
	return ac.sqlConfig
}

const (
	config = "/snlt/config.yml"
)

func NewAppConfig() (AppConfig, error) {
	f, err := os.Open(config)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	var ac appConfig

	err = yaml.NewDecoder(f).Decode(&ac)
	return ac, err
}
