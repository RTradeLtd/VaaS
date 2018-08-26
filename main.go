package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) > 3 || len(os.Args) < 2 {
		log.Fatal("bad args")
	}
	switch os.Args[1] {
	case "api":
		api := InitializeAPI(nil)
		err := api.Router.Run("127.0.0.1:6767")
		if err != nil {
			log.Fatal(err)
		}
	case "distributor":
		if len(os.Args) > 3 || len(os.Args) < 3 {
			fmt.Println("./VaaS distributor <listen-address>")
			log.Fatal("bad args for distributor")
		}
		err := InitializeDistributor(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
	}
}
