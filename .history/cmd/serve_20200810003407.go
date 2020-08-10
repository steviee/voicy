/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run the service on the local node",
	Long: `server starts the tiny web-service to receive voice messages from the running website.
	This is the initial version and only very bare.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := serve()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	serveCmd.PersistentFlags().String("port", "4000", "the port to listen on")
	serveCmd.PersistentFlags().String("bind", "0.0.0.0", "the IP to bind on")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func serve() error {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
  
	// Routes
	e.GET("/", hello)
  
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
  }
  
  // Handler
  func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
  }}
