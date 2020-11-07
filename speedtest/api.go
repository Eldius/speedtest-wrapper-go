package speedtest

import (
	"bytes"
	"encoding/json"
	"os/exec"
)

func Test() (*SpeedtestResult, error) {
	cmd := exec.Command("speedtest", "-f", "json")

	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	var result *SpeedtestResult
	if err := json.Unmarshal(out.Bytes(), &result); err != nil {
		return nil, err
	}
	return result, nil
}
