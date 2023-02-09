package cmd

import (
	"fmt"
	"github.com/Eldius/speedtest-wrapper-go/config"
	"github.com/Eldius/speedtest-wrapper-go/persistence"
	"log"
	"time"

	"github.com/Eldius/speedtest-wrapper-go/speedtest"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run speedtest benchmark",
	Args:  cobra.ExactArgs(0),
	Long: `Run speedtest benchmark. For example:

speedtest-wrapper-go test
`,
	Run: func(_ *cobra.Command, _ []string) {
		repo, err := persistence.NewPersistence(config.GetDBFile())
		if err != nil {
			log.Fatalf("Failed to create persistence object: %v", err)
		}
		if r, err := speedtest.Test(); err != nil {
			fmt.Println("Failed to execute test:", err)
			if *testPublish {
				info := speedtest.SpeedtestResult{
					Timestamp: time.Now(),
				}
				if err := repo.Persist(info); err != nil {
					log.Fatalf("Failed to persist data: %v", err)
				}
			}
		} else {
			fmt.Printf(`---
ISP: %s
server: %s
- download:    %f mbps => megabits/sec (%d bps => bytes/sec)
- upload:      %f mbps => megabits/sec (%d bps => bytes/sec)
- ping:        %f ms
- jitter:      %f ms
- packet loss: %f %%
`,
				r.Isp,
				r.Server.Name,
				r.DownloadInMbps(),
				r.Download.Bandwidth,
				r.UploadInMbps(),
				r.Upload.Bandwidth,
				r.Ping.Latency,
				r.Ping.Jitter,
				r.PacketLoss,
			)

			if *testPublish {
				if err := repo.Persist(*r); err != nil {
					log.Fatalf("Failed to persist data: %v", err)
				}
			}
		}
	},
}

var (
	testPublish *bool
)

func init() {
	rootCmd.AddCommand(testCmd)
	testPublish = testCmd.Flags().BoolP("publish", "p", false, "Publish to MQTT broker")
}
