package main

import (
	"flag"
	"github.com/shantanubansal/ipam/dhcp"
	"log"
)

type cliConfig struct {
	dhcpConfig *string
	port *string
}
const DhcpConfigPath = "dhcp.properties.path"
const ServerPort = "server.port"

func initializeCliFlags() cliConfig {
	var flags cliConfig
	flags.dhcpConfig = flag.String(DhcpConfigPath, "", "a string")
	flags.port = flag.String(ServerPort, "8080", "a string")
	flag.Parse()
	return flags
}


func main()  {
	cliFlags := initializeCliFlags()
	if cliFlags.dhcpConfig == nil || *cliFlags.dhcpConfig == ""{
		log.Fatalf("please provide dhcp properties path. --%v flag missing", DhcpConfigPath)
	}
	dhcp.StartServer(*cliFlags.dhcpConfig, *cliFlags.port)
}