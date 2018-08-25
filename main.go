package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) > 3 || len(os.Args) < 2 {
		log.Fatal("bad args")
	}
	switch os.Args[1] {
	case "api":
		api := InitializeAPI()
		err := api.Router.Run("127.0.0.1:6767")
		if err != nil {
			log.Fatal(err)
		}
	case "distributor":
		err := InitializeDistributor(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
	}
}
