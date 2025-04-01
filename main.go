package main

import (
	"os"

	"github.com/cloudyjeep/nusantara-data/api"
	"github.com/cloudyjeep/nusantara-data/lib"
)

func main() {
	port := lib.Trim(os.Getenv("PORT_ENV"))

	if os.Getenv("GO_ENV") == "development" && port == "" {
		port = "3000"
	}

	api.Init(port)
}
