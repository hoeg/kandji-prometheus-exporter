package collector

func Blueprints(devices []Device) map[string]int {
	r := make(map[string]int)
	for _, device := range devices {
		r[device.BlueprintName] = r[device.BlueprintName] + 1
	}
	return r
}
