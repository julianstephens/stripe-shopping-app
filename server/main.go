package main

import (
	"github.com/julianstephens/stripe-shopping-app/server/router"
	"github.com/julianstephens/stripe-shopping-app/server/utils"
)

func main() {
	// Load server port
	PORT := utils.GetEnvVar("API_PORT")

  // Initialize echo
  e := router.Init()

  // Start the server
  e.Logger.Fatal(e.Start(":" + PORT))
}
