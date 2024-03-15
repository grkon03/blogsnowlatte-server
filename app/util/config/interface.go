package config

type ExecConfig interface {
	DoMiguration() bool
}

type SQLConfig interface {
	User() string
	Pass() string
	Host() string
	Name() string
	Port() uint
}

type AppConfig interface {
	ExecConfig() ExecConfig
	SQLConfig() SQLConfig
}
