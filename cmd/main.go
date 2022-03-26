package main

import (
	"strider-backend-test.com/api"
)

func main() {
	server := api.NewServer()
	server.Run()
}
