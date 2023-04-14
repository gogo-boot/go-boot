package myConfig

import (
	"fmt"
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
	EndPoint     string   `mapstructure:"endPoint"`
}

var MyConfig binder

func LoadConfig() {

	config.WithOptions(config.ParseEnv)

	// only add decoder
	// config.SetDecoder(config.Yaml, yamlv3.Decoder)
	// Or
	config.AddDriver(yamlv3.Driver)

	err := config.LoadFiles("config/config.yml")
	if err != nil {
		panic(err)
	}

	config.BindStruct("", &MyConfig)
	//fmt.Printf("config data: \n %#v\n", config.Data())
	fmt.Printf("config data: \n %v\n", MyConfig)
}
