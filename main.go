package main

import "log"

func main() {
	api := InitializeAPI()
	err := api.Router.Run("127.0.0.1:6767")
	if err != nil {
		log.Fatal(err)
	}
}
