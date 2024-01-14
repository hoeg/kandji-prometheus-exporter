package collector

import "fmt"

func AccumulateVersions(devices []Device) map[string]int {
	r := make(map[string]int)
	for _, device := range devices {
		os := fmt.Sprintf("%s %s", device.Platform, device.OsVersion)
		r[os] = r[os] + 1
	}
	return r
}
