package dhcp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)


const defaultPort = "8080"

func StartServer(configFile, port string) {
	if configFile != "" {
		defaultConfigPath = configFile
	}
	if port == "" {
		port = defaultPort
	}
	http.HandleFunc("/init", Init)
	http.HandleFunc("/ip", IpAllocater)
	err := http.ListenAndServe(fmt.Sprintf(":%v",port), nil)
	if err != nil {
		log.Fatalf("unable to start the server %v", err)
	}
	log.Printf("Listening on port %v", port)
}

func Init(w http.ResponseWriter, req *http.Request) {
	dhcp, err := GetDhcpRanges(defaultConfigPath)
	if err != nil {
		errorResponse(err,w)
	}
	dhcps = dhcp
	successResponse("DHCPs successfully initialized",nil, w)
}


func IpAllocater(w http.ResponseWriter, req *http.Request) {

	inputDhcpInterface := req.URL.Query().Get("dhcp")

	if inputDhcpInterface == "" {
		errorResponse( fmt.Errorf("dhcp interface name can only have one value"), w)
	}
		ip, err := allocateIp(inputDhcpInterface)
	if err != nil {
		errorResponse(err, w)
	}
	successResponse("ip allocated", ip, w)
}

func allocateIp(inputDhcpInterface string) (string, error){
	dhcp := dhcps[inputDhcpInterface]
	numberOfAvailableIps := len(dhcp.AvailableIps)
	if numberOfAvailableIps < 1 {
		return "", fmt.Errorf("no ip available for allocation for requested dhcp interface")
	}
	allocatedIp := dhcp.AvailableIps[numberOfAvailableIps-1]
	dhcp.AllocatedIps = append(dhcp.AllocatedIps, allocatedIp)
	dhcp.AvailableIps = dhcp.AvailableIps[0 : numberOfAvailableIps-1]
	dhcps[inputDhcpInterface] = dhcp
	return allocatedIp, nil
}

func errorResponse( err error, w http.ResponseWriter) {
	e := response( HttpResponse{
		Code: 500,
		Msg:  fmt.Sprintf("unble to process request due to %v,", err),
	}, w)
	if e != nil {
		log.Printf("unable to return error due to %v", e)
	}
}

func successResponse(msg string,data interface{}, w http.ResponseWriter) {
	err :=  response( HttpResponse{
		Code: 200,
		Msg:  msg,
		Data: data,
	}, w)
	if err != nil {
		log.Printf("unable to return error due to %v", err)
	}
}

func response(httpRes HttpResponse, w http.ResponseWriter) error {
	bytes, err := json.Marshal(httpRes)
	if err != nil {
		return err
	}
	_, err = w.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}
type HttpResponse struct {
	Code int
	Msg string
	Data interface{}
}




