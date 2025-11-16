package main

import (
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"strings"
)

func rpcClientSend(command string, cfg *Config) {

	// connect to RPC server
	address := fmt.Sprintf("%s:%d", cfg.server, cfg.port)
	fmt.Printf("Connecting to RPC server at %s...\n", address)
	client, err := rpc.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// send command to server
	cmd := strings.ToLower(command)
	switch cmd {

	case "clear":
		sendMessage(client, "Clear")
	case "display":
		sendMessage(client, "Display")
	case "system":
		sendMessage(client, "System")
	case "critical":
		sendMessage(client, "Critical")
	case "read":
		flags := sendMessage(client, "Read")
		log.Printf("Previous ThreadExecutionState flags: 0x%X", flags)
	case "shutdown":
		sendMessage(client, "Shutdown")
	default:
		flag.Usage()
		return
	}
}

// sendMessage sends an RPC call to the server.
//
// Parameters:
//
//	client - the RPC client used to communicate with the server
//	method - the method name (string) to call on the SleepControl service
func sendMessage(client *rpc.Client, method string) uint32 {
	var args struct{}
	var reply ExecStateReply
	err := client.Call("ExecStateManager."+method, &args, &reply)
	if err != nil {
		log.Fatalf("RPC error in %s: %v", method, err)
	}
	log.Printf("Successfully sent %s RPC", method)
	return reply.Flags
}
