package main

import (
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
)

func rpcClientSend(command string, cfg *Config) {

	// connect to RPC server
	address := fmt.Sprintf("%s:%d", cfg.address, cfg.port)
	fmt.Printf("Connecting to RPC server at %s (%s) ...\n", address, cfg.network)
	client, err := rpc.Dial(cfg.network, address)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close() //nolint:errcheck

	// send command to server
	cmd := strings.ToLower(command)
	switch cmd {

	case "clear":
		sendMessage(client, "Clear", struct{}{})
	case "display":
		sendMessage(client, "Display", struct{}{})
	case "system":
		sendMessage(client, "System", struct{}{})
	case "critical":
		sendMessage(client, "Critical", struct{}{})
	case "read":
		reply := sendMessage(client, "Read", struct{}{})
		log.Printf("Previous ThreadExecutionState flags: 0x%X", reply.Flags)
		log.Printf("Registered processes: %v", reply.Processes)
	case "register":
		sendMessage(client, "Register", ExecStateRequest{Process: os.Getpid()})
	case "unregister":
		sendMessage(client, "Unregister", ExecStateRequest{Process: os.Getpid()})
	case "shutdown":
		sendMessage(client, "Shutdown", struct{}{})
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
func sendMessage(client *rpc.Client, command string, args any) ExecStateReply {
	var reply ExecStateReply
	err := client.Call("ExecStateManager."+command, args, &reply)
	if err != nil {
		log.Fatalf("RPC error in %s: %v", command, err)
	}
	log.Printf("Successfully sent %s RPC", command)
	return reply
}
