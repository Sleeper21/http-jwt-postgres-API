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
	// c.VueSPA = os.Getenv("VUE_SPA_HOST")

	// c.PingIdClientId = os.Getenv("PING_CLIENT_ID")
	// c.PingIdClientSecret = os.Getenv("PING_CLIENT_SECRET")
	// c.PingIdUrl = os.Getenv("PING_URL")
	// c.PingIdTokenUrl = os.Getenv("PING_TOKEN_URL")
	// c.PingIdCallbackUrl = os.Getenv("PING_CALLBACK_URL")
	// c.PingJWKSUrl = os.Getenv("PING_JWKS_URL")
	// c.PingIdUserInfoUrl = os.Getenv("PING_USERINFO_URL")

	// if groupAccess := os.Getenv("USER_GROUP_ACCESS"); groupAccess != "" {
	// 	c.UserGroupAccess = strings.Split(groupAccess, ",")
	// }

	return c, nil
}
