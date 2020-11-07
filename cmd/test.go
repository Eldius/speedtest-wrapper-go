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
	"fmt"

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
	Run: func(cmd *cobra.Command, args []string) {
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
		}

	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
