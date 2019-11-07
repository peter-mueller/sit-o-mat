package admincli

import (
	"log"
	"net/http"
)

type AssignWorkplaces struct {
}

func (r AssignWorkplaces) Execute() {
	req, _ := http.NewRequest(
		"GET",
		"http://"+config.Host+"/sitomat",
		nil)
	req.SetBasicAuth(config.AdminUsername, config.AdminPassword)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	res.Write(log.Writer())
}
