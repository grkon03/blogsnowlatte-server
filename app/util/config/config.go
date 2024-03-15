package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ExecConfig_ struct {
	DoMiguration_ bool `yaml:"do_miguration"`
}

func (ec ExecConfig_) DoMiguration() bool {
	return ec.DoMiguration_
}

type SQLConfig_ struct {
	User_ string `yaml:"user"`
	Pass_ string `yaml:"pass"`
	Host_ string `yaml:"host"`
	Name_ string `yaml:"name"`
	Port_ uint   `yaml:"port"`
}

func (sc SQLConfig_) User() string {
	return sc.User_
}
func (sc SQLConfig_) Pass() string {
	return sc.Pass_
}
func (sc SQLConfig_) Host() string {
	return sc.Host_
}
func (sc SQLConfig_) Name() string {
	return sc.Name_
}
func (sc SQLConfig_) Port() uint {
	return sc.Port_
}

type AppConfig_ struct {
	Ec ExecConfig_ `yaml:"exec_config"`
	Sc SQLConfig_  `yaml:"sql_config"`
}

func (ac AppConfig_) ExecConfig() ExecConfig {
	return ac.Ec
}

func (ac AppConfig_) SQLConfig() SQLConfig {
	return ac.Sc
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

	var ac AppConfig_

	err = yaml.NewDecoder(f).Decode(&ac)
	if err != nil {
		panic(err)
	}

	return ac, nil
}
