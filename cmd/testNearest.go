package cmd

import (
	"fmt"
	"github.com/Eldius/speedtest-wrapper-go/config"
	"github.com/Eldius/speedtest-wrapper-go/persistence"
	"log"

	"github.com/Eldius/speedtest-wrapper-go/speedtest"
	"github.com/spf13/cobra"
)

// testNearestCmd represents the testNearest command
var testNearestCmd = &cobra.Command{
	Use:   "testNearest",
	Short: "Test internet speed in the nearest servers",
	Long: `Test internet speed in the nearest servers. For example:

	speedtest-wrapper-go testNearest
`,
	Run: func(_ *cobra.Command, _ []string) {
		cfg := config.AppConfig()
		repo, err := persistence.NewPersistence(cfg.DBFile)
		if err != nil {
			log.Fatalf("Failed to create persistence object: %v", err)
		}

		servers, err := speedtest.ListServers()
		if err != nil {
			log.Fatalf("Failed to list nearest servers: %s", err.Error())
		}
		for _, s := range servers.Servers {
			if r, err := speedtest.TestForServer(s); err != nil {
				fmt.Printf("Failed to execute test: %s\n", err.Error())
				if *nearestPublish {
					if err := repo.Persist(*r); err != nil {
						log.Fatalf("Failed to persist data: %v", err)
					}
				}
			} else {
				fmt.Printf(r.ToString())
				if *nearestPublish {
					if err := repo.Persist(*r); err != nil {
						log.Fatalf("Failed to persist data: %v", err)
					}
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
