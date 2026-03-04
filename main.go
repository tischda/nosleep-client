package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const DEFAULT_PORT = 9001

// https://goreleaser.com/cookbooks/using-main.version/
var (
	name    string
	version string
	date    string
	commit  string
)

// flags
type Config struct {
	network string
	address string
	port    int
	help    bool
	version bool
	pid     int
}

func initFlags() *Config {
	cfg := &Config{}
	flag.StringVar(&cfg.network, "n", "tcp", "")
	flag.StringVar(&cfg.network, "network", "tcp", "Network type (tcp, tcp4, tcp6, unix, etc.)")
	flag.StringVar(&cfg.address, "a", "127.0.0.1", "")
	flag.StringVar(&cfg.address, "address", "127.0.0.1", "RPC server")
	flag.IntVar(&cfg.port, "p", DEFAULT_PORT, "")
	flag.IntVar(&cfg.port, "port", DEFAULT_PORT, "RPC server listening port")
	flag.BoolVar(&cfg.help, "?", false, "")
	flag.BoolVar(&cfg.help, "help", false, "displays this help message")
	flag.BoolVar(&cfg.version, "v", false, "")
	flag.BoolVar(&cfg.version, "version", false, "print version and exit")
	return cfg
}

// Make sure you use the same structs on both client and server side for the RPC calls to work correctly.
type ExecStateRequest struct {
	Process int
}

type ExecStateReply struct {
	Flags     uint32
	Processes []int
}

func main() {
	log.SetFlags(0)
	cfg := initFlags()

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: "+name+` [OPTIONS] <COMMAND> 

Calls the NoSleep RPC server on ADDRESS:PORT (default: 127.0.0.1:`+fmt.Sprintf("%d", DEFAULT_PORT)+`).
You can manage the server using RPC calls to control thread execution states.

COMMANDS:

   Clear, Display, System, Critical, Read, Shutdown
   Register, Unregister --pid <PID>

OPTIONS:

  -n, --network string
          Network type: tcp, tcp4, tcp6, unix or unixpacket (default "tcp")
  -a, --address string
          Bind address (default 127.0.0.1)
  -p, --port int
          RPC server listening port (default 9001)
  -?, --help
        displays this help message
  -v, --version
        print version and exit

EXAMPLES:`)

		fmt.Fprintln(os.Stderr, "\n  "+name+` --port 9015 display

  will set ThreadExecutionState to (ES_CONTINUOUS | ES_SYSTEM_REQUIRED | ES_DISPLAY_REQUIRED)`)
	}

	flag.Parse()

	if flag.Arg(0) == "version" || cfg.version {
		fmt.Printf("%s %s, built on %s (commit: %s)\n", name, version, date, commit)
		return
	}

	if cfg.help {
		flag.Usage()
		return
	}

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	// Handle Register/Unregister subcommands with --pid flag
	command := strings.ToLower(flag.Arg(0))
	if command == "register" || command == "unregister" {
		subFlags := flag.NewFlagSet(command, flag.ContinueOnError)
		subFlags.IntVar(&cfg.pid, "pid", 0, "PID of the process to register/unregister")
		if err := subFlags.Parse(flag.Args()[1:]); err != nil {
			os.Exit(1)
		}
		if cfg.pid <= 0 {
			log.Fatalf("flag --pid required to be greater than 0 for %s", command)
		}
	}

	// send command to server
	rpcClientSend(flag.Arg(0), cfg)
}
