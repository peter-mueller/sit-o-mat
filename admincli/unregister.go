package admincli

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type Unregister struct {
	Name string
}

func ParseUnregister(arguments []string) Unregister {
	unregister := Unregister{}
	var unregisterCmd = flag.NewFlagSet("unregister", flag.ExitOnError)
	unregisterCmd.StringVar(&unregister.Name, "name", "", "username of user to unregister")
	unregisterCmd.Parse(arguments)

	if unregister.Name == "" {
		log.Println("missing user name")
		unregisterCmd.Usage()
		os.Exit(1)
	}
	return unregister
}

func (r Unregister) Execute() {
	req, _ := http.NewRequest(
		"DELETE",
		"http://"+config.Host+"/user/"+r.Name,
		nil)
	req.SetBasicAuth(config.AdminUsername, config.AdminPassword)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	res.Write(log.Writer())
}
