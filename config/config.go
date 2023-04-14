package myConfig

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
)

type binder struct {
	Server server `mapstructure:"server"`
	Oauth2 oauth2 `mapstructure:"oauth2"`
}

type server struct {
	PortNumber string `mapstructure:"portNumber"`
}

type oauth2 struct {
	RedirectUrl  string   `mapstructure:"redirectUrl"`
	ClientId     string   `mapstructure:"clientId"`
	ClientSecret string   `mapstructure:"clientSecret"`
	Scopes       []string `mapstructure:"scopes"`
	Tenant       string   `mapstructure:"tenant"`
}

var AppConfig binder

func init() {

	config.WithOptions(config.ParseEnv)
	config.AddDriver(yamlv3.Driver)

	err := config.LoadFiles("config/config.yml")
	if err != nil {
		panic(err)
	}

	config.BindStruct("", &AppConfig)
	//fmt.Printf("config data: \n %v\n", AppConfig)
}
