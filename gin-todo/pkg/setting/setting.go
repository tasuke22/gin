package setting

import (
	"log"

	"github.com/go-ini/ini"
)

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

var DatabaseSetting = &Database{}

var cfg *ini.File

func Setup(iniPath string) {
	var err error
	cfg, err = ini.Load(iniPath)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse %s: %v", iniPath, err)
	}

	mapTo("database", DatabaseSetting)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
