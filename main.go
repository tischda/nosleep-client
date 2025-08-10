package main

import (
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
)

const PROG_NAME string = "nosleep-client"
const DEFAULT_PORT = 9001

var version string

var flagHelp = flag.Bool("help", false, "displays this help message")
var flagPort = flag.Int("port", DEFAULT_PORT, "RPC server listening port")
var flagVersion = flag.Bool("version", false, "print version and exit")

type ExecStateReply struct {
	Flags uint32
}

func init() {
	flag.BoolVar(flagHelp, "h", false, "")
	flag.IntVar(flagPort, "p", DEFAULT_PORT, "")
	flag.BoolVar(flagVersion, "v", false, "")
}

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: "+PROG_NAME+` [--port <port>] <COMMAND> | --version | --help

Calls the NoSleep RPC server on 127.0.0.1:`+fmt.Sprintf("%d", DEFAULT_PORT)+`.
You can manage the server using RPC calls to control thread execution states.

COMMANDS:

   Clear, Display, System, Critical, Read and Shutdown.

OPTIONS:

  -h, -help
        displays this help message
  -p, -port int
        RPC server listening port (default 9001)
  -v, -version
        print version and exit

EXAMPLES:`)

		fmt.Fprintln(os.Stderr, "\n  "+PROG_NAME+` --port 9015 display

  will set ThreadExecutionState to (ES_CONTINUOUS | ES_SYSTEM_REQUIRED | ES_DISPLAY_REQUIRED)`)
	}

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
	cmd := strings.ToLower(flag.Arg(0))
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
