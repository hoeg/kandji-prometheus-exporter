package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const kandjiURL string = "kandji.io"

func ListDevices() ([]Device, error) {
	urlFmt := "https://%s.api.kandji.io/api/v1/devices"

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(urlFmt, kandjiURL), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer {api_token}")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return unmarshalDevices(body)
}

func unmarshalDevices(body []byte) ([]Device, error) {
	devices := []Device{}
	err := json.Unmarshal(body, &devices)
	return devices, err
}

type Device struct {
	DeviceID                   string    `json:"device_id"`
	DeviceName                 string    `json:"device_name"`
	Model                      string    `json:"model"`
	SerialNumber               string    `json:"serial_number"`
	Platform                   string    `json:"platform"`
	OsVersion                  string    `json:"os_version"`
	SupplementalBuildVersion   string    `json:"supplemental_build_version"`
	SupplementalOsVersionExtra string    `json:"supplemental_os_version_extra"`
	LastCheckIn                time.Time `json:"last_check_in"`
	User                       string    `json:"user"`
	AssetTag                   string    `json:"asset_tag"`
	BlueprintID                string    `json:"blueprint_id"`
	MdmEnabled                 bool      `json:"mdm_enabled"`
	AgentInstalled             bool      `json:"agent_installed"`
	IsMissing                  bool      `json:"is_missing"`
	IsRemoved                  bool      `json:"is_removed"`
	AgentVersion               string    `json:"agent_version"`
	FirstEnrollment            string    `json:"first_enrollment"`
	LastEnrollment             string    `json:"last_enrollment"`
	BlueprintName              string    `json:"blueprint_name"`
	LostModeStatus             string    `json:"lost_mode_status"`
}
