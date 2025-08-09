package main

import (
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"os"
)

const PROG_NAME string = "nosleep-client"
const DEFAULT_PORT = 9001

var version string

var flagHelp = flag.Bool("help", false, "displays this help message")
var flagPort = flag.Int("port", DEFAULT_PORT, "RPC server listening port")
var flagVersion = flag.Bool("version", false, "print version and exit")

func init() {
	flag.BoolVar(flagHelp, "h", false, "")
	flag.IntVar(flagPort, "p", DEFAULT_PORT, "")
	flag.BoolVar(flagVersion, "v", false, "")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" || *flagVersion {
		fmt.Printf("%s version %s\n", PROG_NAME, version)
		return
	}

	if *flagHelp {
		flag.Usage()
		return
	}

	if flag.NArg() == 0 || flag.NArg() > 1 {
		flag.Usage()
		os.Exit(1)
	}

	// Create RPC client
	address := fmt.Sprintf("127.0.0.1:%d", *flagPort)
	client, err := rpc.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Send command to server
	switch flag.Arg(0) {
	case "sleep":
		sendMessage(client, "Sleep")
	case "display":
		sendMessage(client, "Display")
	case "system":
		sendMessage(client, "System")
	case "critical":
		sendMessage(client, "Critical")
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
func sendMessage(client *rpc.Client, method string) {
	var args, reply struct{}
	err := client.Call("SleepControl."+method, &args, &reply)
	if err != nil {
		log.Fatalf("RPC error in %s: %v", method, err)
	}
	log.Printf("Successfully sent %s RPC", method)
}
