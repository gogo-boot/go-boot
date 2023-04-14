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
	PortNumber uint16 `mapstructure:"portNumber"`
}

type oauth2 struct {
	RedirectUrl  string   `mapstructure:"redirectUrl"`
	ClientId     string   `mapstructure:"clientId"`
	ClientSecret string   `mapstructure:"clientSecret"`
	Scopes       []string `mapstructure:"scopes"`
	Tenant       string   `mapstructure:"tenant"`
}

var MyConfig binder

func init() {

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
	//fmt.Printf("config data: \n %v\n", MyConfig)
}
