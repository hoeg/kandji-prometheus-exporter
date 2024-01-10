package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const kandjiURL string = "kandji.io"

type Collector struct {
	kandjiURL string
	token     string
}

func NewCollector(kandjiURL, token string) *Collector {
	return &Collector{
		kandjiURL: kandjiURL,
		token:     token,
	}
}

func (c *Collector) ListDevices() ([]Device, error) {
	offset := 0

	devices := []Device{}
	for {
		d, err := c.deviceChunk(offset)
		if err != nil {
			return nil, err
		}
		if len(d) == 0 {
			break
		}
		devices = append(devices, d...)
		offset += 300
	}
	return devices, nil
}

func (c *Collector) deviceChunk(offset int) ([]Device, error) {
	urlFmt := "%s/api/v1/devices?limit=300&offset=%d"

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(urlFmt, c.kandjiURL, offset), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if len(body) == 0 {
		return nil, nil
	}
	return unmarshalDevices(body)
}

func unmarshalDevices(body []byte) ([]Device, error) {
	var devices []Device
	err := json.Unmarshal(body, &devices)
	return devices, err
}

type Device struct {
	DeviceID                   string    `json:"device_id,omitempty"`
	DeviceName                 string    `json:"device_name,omitempty"`
	Model                      string    `json:"model,omitempty"`
	SerialNumber               string    `json:"serial_number,omitempty"`
	Platform                   string    `json:"platform,omitempty"`
	OsVersion                  string    `json:"os_version,omitempty"`
	SupplementalBuildVersion   string    `json:"supplemental_build_version,omitempty"`
	SupplementalOsVersionExtra string    `json:"supplemental_os_version_extra,omitempty"`
	LastCheckIn                time.Time `json:"last_check_in,omitempty"`
	User                       string    `json:"user,omitempty"`
	AssetTag                   string    `json:"asset_tag,omitempty"`
	BlueprintID                string    `json:"blueprint_id,omitempty"`
	MdmEnabled                 bool      `json:"mdm_enabled,omitempty"`
	AgentInstalled             bool      `json:"agent_installed,omitempty"`
	IsMissing                  bool      `json:"is_missing,omitempty"`
	IsRemoved                  bool      `json:"is_removed,omitempty"`
	AgentVersion               string    `json:"agent_version,omitempty"`
	FirstEnrollment            string    `json:"first_enrollment,omitempty"`
	LastEnrollment             string    `json:"last_enrollment,omitempty"`
	BlueprintName              string    `json:"blueprint_name,omitempty"`
	LostModeStatus             string    `json:"lost_mode_status,omitempty"`
}
