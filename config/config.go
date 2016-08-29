package config

import (
	"io/ioutil"
	"path/filepath"

	"github.com/go-yaml/yaml"
)

type Configuration struct {
	Certificate   Certificate `yaml:"certificate"`
	Database      string
	Storage       Storage       `yaml:"storage"`
	License       License       `yaml:"license"`
	LsdServer     ServerInfo    `yaml:"lsd"`
	LcpServer     ServerInfo    `yaml:"lcp"`
	Static        Static        `yaml:"static"`
	LicenseStatus LicenseStatus `yaml:"license_status"`
	Localization  Localization  `yaml:"localization"`
}

type ServerInfo struct {
	Host          string `yaml:"host,omitempty"`
	Port          int    `yaml:"port,omitempty"`
	ReadOnly      bool   `yaml:"readonly,omitempty"`
	PublicBaseUrl string `yaml:"public_base_url,omitempty"`
	Database      string `yaml:"database,omitempty"`
}

type Certificate struct {
	Cert       string `yaml:"cert"`
	PrivateKey string `yaml:"private_key"`
}

type FileSystem struct {
	Directory string `yaml:"directory"`
}

type Static struct {
	Directory string `yaml:"directory"`
}

type Storage struct {
	FileSystem FileSystem `yaml:"filesystem"`
	AccessId   string     `yaml:"access_id"`
	DisableSSL bool       `yaml:"disable_ssl"`
	PathStyle  bool       `yaml:"path_style"`
	Mode       string
	Secret     string
	Endpoint   string
	Bucket     string
	Region     string
	Token      string
}

type License struct {
	Links map[string]string `yaml:"links"`
}

type LicenseStatus struct {
	Renew       bool `yaml:"renew"`
	Register    bool `yaml:"register"`
	Return      bool `yaml:"return"`
	RentingDays int  `yaml:"renting_days" "default 0"`
	RenewDays   int  `yaml:"renew_days" "default 0"`
}

type Localization struct {
	Languages       []string `yaml:"languages"`
	Folder          string   `yaml:"folder"`
	DefaultLanguage string   `yaml:"default_language"`
}

var Config Configuration

func ReadConfig(configFileName string) {
	filename, _ := filepath.Abs(configFileName)
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic("Can't read config file: " + configFileName)
	}

	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic("Can't unmarshal config. " + err.Error())
	}
}
