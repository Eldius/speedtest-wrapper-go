package speedtest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
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

func TestForServer(s Server) (*SpeedtestResult, error) {
	cmd := exec.Command("speedtest", "-f", "json", "-s", strconv.Itoa(s.ID))

	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("Failed to execute test for server '%s': %s => execution result: %s", s.Name, err.Error(), string(out.Bytes()))
	}
	var result *SpeedtestResult
	if err := json.Unmarshal(out.Bytes(), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func ListServers() (*ServerList, error) {
	cmd := exec.Command("speedtest", "-L", "-f", "json")

	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	var result *ServerList
	if err := json.Unmarshal(out.Bytes(), &result); err != nil {
		return nil, err
	}
	return result, nil
}
