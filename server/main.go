package main

import (
	"github.com/julianstephens/stripe-shopping-app-tutorial/server/router"
	"github.com/julianstephens/stripe-shopping-app-tutorial/server/utils"
)

func main() {
	// Load server port
	PORT := utils.GetEnvVar("API_PORT")

  // Initialize echo
  e := router.Init()

  // Start the server
  e.Logger.Fatal(e.Start(":" + PORT))
}
