Target:  A very generic DHCP Server and Ip Calculator


# ipam

The basic implementation of Ip Allocation


To Run Server, run the main file with argument --dhcp.properties.path 

How to run from the ipam root dir:

`cd ipam`

`go run ./main.go --dhcp.properties.path ./dhcp/dhcp_config.properties
`

you can provide --server.port if want to alter the port of the server (default 8080)

` go run ./main.go --dhcp.properties.path ./dhcp/dhcp_config.properties --server.port 8081
`

Run:
It will load the configuration

`http://localhost:8080/init`

Then to get the ip run

`http://localhost:8080/ip/add?dhcp=<dhcp_interface>`


Then to free the ip 

`http://localhost:8080/ip/remove?dhcp=<dhcp_interface>addreess=<ipAddress>`

eg:

`http://localhost:8080/ip?dhcp=eth1`
