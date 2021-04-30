package dhcp

import (
	"path"
	"runtime"
)

type Dhcp struct {
	AvailableIps []string
	AllocatedIps []string
}

var dhcps map[string]Dhcp

var defaultConfigPath = "./dhcp_config.properties"

func init () {
	dhcps = make(map[string]Dhcp)
	_, fileName, _, _ := runtime.Caller(1)
	defaultConfigPath = path.Join(path.Dir(fileName), "dhcp_config.properties")
}

