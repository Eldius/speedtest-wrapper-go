package cmd

import (
	"fmt"

	"github.com/Eldius/speedtest-wrapper-go/config"
	"github.com/Eldius/speedtest-wrapper-go/mqttclient"
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
		if r, err := speedtest.Test(); err != nil {
			fmt.Println("Failed to execute test")
			fmt.Println(err.Error())
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
				mqttclient.SendTestResult(r.CreateSummary(), config.AppConfig().MQTT)
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
