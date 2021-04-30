package dhcp

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readConfigFile(path string) ([]string,error){
	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil,err
	}
	fileDataStr := string(fileData)
	fileDataArray := strings.Split(fileDataStr,"\n")
	return fileDataArray, nil
}

const (
	dhcpRangeKey = "dhcp-range"
)


func GetDhcpRanges(path string) (map[string]Dhcp, error){
	fileLines, err := readConfigFile(path)
	if err != nil {
		return nil, err
	}
	dhcps := make(map[string]Dhcp)
	for _,line := range fileLines{
		if strings.Contains(line, dhcpRangeKey) {
			dhcp, ips, err := dhcpInfo(line)
			if err != nil {
				return nil, err
			}
			dhcps[dhcp] = ips
		}
	}
	return dhcps, nil
}

func dhcpInfo(line string) (string,Dhcp, error) {
	values := strings.Split(line, "=")
	ipInfo := strings.Split(values[1], ",")
	dhcpName := ipInfo[0]
	startIp := ipInfo[1]
	endIp := ipInfo[2]
	ipParts := strings.Split(startIp, ".")
	ipStart, err := strconv.ParseInt(ipParts[3], 10, 64)
	if err != nil {
		return "", Dhcp{}, fmt.Errorf("unable to parse ip parts due to %v", err)
	}
	ipEnd, err := strconv.ParseInt(strings.Split(endIp, ".")[3], 10, 64)
	if err != nil {
		return "", Dhcp{}, fmt.Errorf("unable to parse ip parts due to %v", err)
	}
	availableIps := make([]string, 0)
	for i := ipStart; i <= ipEnd; i++ {
		availableIps = append(availableIps, fmt.Sprintf("%v.%v.%v.%v", ipParts[0], ipParts[1], ipParts[2], i))
	}
	return dhcpName,  Dhcp{
		AvailableIps: availableIps,
		AllocatedIps: make([]string, 0),
	}, nil
}
