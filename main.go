package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2" // Importing the third-party CLI library
)

func main() {
	// Creating a new CLI application instance
	app := &cli.App{
		Name:  "HealthChecker", // Setting the name of the CLI application
		Usage: "A tiny tool that checks whether a website is running or is down", // Setting the usage description
		Flags: []cli.Flag{ // Defining command line flags
			&cli.StringFlag{
				Name:     "domain",         // Flag for domain name to check
				Aliases:  []string{"d"},    // Alias for the domain flag
				Usage:    "Domain name to check.", // Description of the domain flag
				Required: true,             // Flag is required
			},
			&cli.StringFlag{
				Name:     "port",           // Flag for port number to check
				Aliases:  []string{"p"},    // Alias for the port flag
				Usage:    "Port number to check.", // Description of the port flag
				Required: false,            // Flag is optional
			},
		},
		Action: func(c *cli.Context) error { // Action to execute when CLI command is run
			port := c.String("port") // Retrieve the value of the port flag
			if port == "" {          // If port flag is not provided, default to port 80
				port = "80"
			}
			status := Check(c.String("domain"), port) // Call function to check status of domain and port
			fmt.Println(status) // Print the status of the domain and port
			return nil // Return nil to indicate success
		},
	}

	// Run the CLI application with the command line arguments
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err) // Log and terminate if there's an error running the CLI application
	}
}
