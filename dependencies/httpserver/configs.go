package httpserver

import (
	"fmt"
	"os"
)

type ServerConfigs struct {
	ApiHost string
	Host    string
	Port    string
}

func LoadServerConfigs() (ServerConfigs, error) {

	c := ServerConfigs{}
	c.Host = os.Getenv("SERVER_HOST")
	c.Port = os.Getenv("SERVER_PORT")
	apiHost := fmt.Sprintf("%s:%s", c.Host, c.Port)
	c.ApiHost = apiHost

	fmt.Println("Api host ", apiHost)

	return c, nil
}
