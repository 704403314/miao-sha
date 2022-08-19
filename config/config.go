package config

import (
	"time"
)
type Config struct {
	Mysql         []MysqlConfig `yaml:"dataSource"` // mysql configuration
	Redis         []RedisConfig // redis configuration
}

type RedisConfig struct {
	Host            string        `yaml:"host"`
	Password        string        `yaml:"password"`
	MaxIdle         int           `yaml:"maxIdle"`   // Connection pool maximum idle connections
	MaxActive       int           `yaml:"maxActive"` // Connection pool maximum active connections
	IdleTimeout     time.Duration `yaml:"idleTimeout"`
	Wait            bool          `yaml:"wait"`            //wait for connection
	MaxConnLifetime time.Duration `yaml:"maxConnLifetime"` //Active time
	//Debug           bool          `yaml:"debug"`           //debug open
}

type MysqlConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Charset      string `yaml:"charset"`
	MaxIdleConns int `yaml:"maxIdleConns"`
	MaxOpenConns int `yaml:"MaxOpenConns"`
	//SqlLog       bool   `yaml:"sql_log"`
}