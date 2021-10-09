/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"log"

	"github.com/Eldius/speedtest-wrapper-go/config"
	"github.com/Eldius/speedtest-wrapper-go/mqttclient"
	"github.com/Eldius/speedtest-wrapper-go/speedtest"
	"github.com/spf13/cobra"
)

// testNearestCmd represents the testNearest command
var testNearestCmd = &cobra.Command{
	Use:   "testNearest",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(_ *cobra.Command, _ []string) {
		servers, err := speedtest.ListServers()
		if err != nil {
			log.Fatalf("Failed to list nearest servers: %s", err.Error())
		}
		for _, s := range servers.Servers {
			if r, err := speedtest.TestForServer(s); err != nil {
				fmt.Printf("Failed to execute test: %s\n", err.Error())
				if *testPublish {
					mqttclient.SendTestResult(speedtest.CreateErrorSummary(err, s), config.AppConfig().MQTT)
				}
			} else {
				fmt.Printf(r.ToString())
				if *testPublish {
					mqttclient.SendTestResult(r.CreateSummary(), config.AppConfig().MQTT)
				}
			}
		}

	},
}

var (
	nearestPublish *bool
)

func init() {
	rootCmd.AddCommand(testNearestCmd)
	nearestPublish = testNearestCmd.Flags().BoolP("publish", "p", false, "Publish to MQTT broker")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testNearestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testNearestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
