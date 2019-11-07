package main

import (
	"log"
	"os"

	"github.com/peter-mueller/sit-o-mat/admincli"
)

func init() {
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal(
			`use one of the commands 
  - sitotmatadmin register
  - ...
			`)
	}

	command := os.Args[1]

	switch command {
	case "register":
		register := admincli.ParseRegister(os.Args[2:])
		register.Execute()
	case "unregister":
		unregister := admincli.ParseUnregister(os.Args[2:])
		unregister.Execute()
	case "addworkplace":
		addworkplace := admincli.ParseAddWorkplace(os.Args[2:])
		addworkplace.Execute()
	case "assignworkplaces":
		admincli.AssignWorkplaces{}.Execute()
	}
}
