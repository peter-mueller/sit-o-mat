package admincli

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
)

var client = http.DefaultClient

var config struct {
	AdminUsername string `required:"true"`
	AdminPassword string `required:"true"`
	Host          string `required:"true"`
}

func init() {
	log.SetFlags(log.Lshortfile)

	err := envconfig.Process("SITOMAT", &config)
	if err != nil {
		message := fmt.Errorf("missing configuration enviroment variable: %v", err)
		log.Fatal(message)
	}
}
