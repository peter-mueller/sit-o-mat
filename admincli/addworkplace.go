package admincli

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/peter-mueller/sit-o-mat/workplace"
)

type AddWorkplace struct {
	workplace.Workplace
}

func ParseAddWorkplace(arguments []string) AddWorkplace {
	addworkplace := AddWorkplace{}
	var addworkplaceCmd = flag.NewFlagSet("addworkplace", flag.ExitOnError)
	addworkplaceCmd.StringVar(&addworkplace.Name, "name", "", "unique name of the workplace")
	addworkplaceCmd.StringVar(&addworkplace.Location, "location", "unknown", "location (e.g. a room) to find the workplace")
	addworkplaceCmd.UintVar(&addworkplace.Ranking, "ranking", 1, "ranking for the workplace")

	addworkplaceCmd.Parse(arguments)

	if addworkplace.Name == "" {
		log.Println("missing workplace name")
		addworkplaceCmd.Usage()
		os.Exit(1)
	}
	return addworkplace
}

func (r AddWorkplace) Execute() {
	body, err := json.Marshal(r.Workplace)
	if err != nil {
		log.Fatal(err)
	}

	req, _ := http.NewRequest(
		"POST",
		"http://"+config.Host+"/workplace",
		bytes.NewBuffer(body))
	req.SetBasicAuth(config.AdminUsername, config.AdminPassword)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	res.Write(log.Writer())
}
