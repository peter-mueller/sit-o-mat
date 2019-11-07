package admincli

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/peter-mueller/sit-o-mat/user"
)

type Register struct {
	Name string
}

func ParseRegister(arguments []string) Register {
	register := Register{}
	var registerCmd = flag.NewFlagSet("register", flag.ExitOnError)
	registerCmd.StringVar(&register.Name, "name", "", "username for the user to create")
	registerCmd.Parse(arguments)

	if register.Name == "" {
		log.Println("missing user name")
		registerCmd.Usage()
		os.Exit(1)
	}
	return register
}

func (r Register) Execute() {
	user := user.User{
		Name: r.Name,
	}
	body, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	req, _ := http.NewRequest(
		"POST",
		"http://"+config.Host+"/user",
		bytes.NewBuffer(body))

	req.SetBasicAuth(config.AdminUsername, config.AdminPassword)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	res.Write(log.Writer())
}
