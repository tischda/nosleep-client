package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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
	server  string
	port    int
	help    bool
	version bool
}

func initFlags() *Config {
	cfg := &Config{}
	flag.StringVar(&cfg.server, "s", "127.0.0.1", "")
	flag.StringVar(&cfg.server, "server", "127.0.0.1", "RPC server")
	flag.IntVar(&cfg.port, "p", DEFAULT_PORT, "")
	flag.IntVar(&cfg.port, "port", DEFAULT_PORT, "RPC server listening port")
	flag.BoolVar(&cfg.help, "?", false, "")
	flag.BoolVar(&cfg.help, "help", false, "displays this help message")
	flag.BoolVar(&cfg.version, "v", false, "")
	flag.BoolVar(&cfg.version, "version", false, "print version and exit")
	return cfg
}

type ExecStateReply struct {
	Flags uint32
}

func main() {
	log.SetFlags(0)
	cfg := initFlags()

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: "+name+` [--server <address>] [--port <port>] <COMMAND> 

Calls the NoSleep RPC server on SERVER:PORT (default: 127.0.0.1:`+fmt.Sprintf("%d", DEFAULT_PORT)+`).
You can manage the server using RPC calls to control thread execution states.

COMMANDS:

   Clear, Display, System, Critical, Read and Shutdown.

OPTIONS:

  -s, --server
        RPC server (default 127.0.0.1)
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

	if flag.NArg() == 0 || flag.NArg() > 1 {
		flag.Usage()
		os.Exit(1)
	}

	// send command to server
	rpcClientSend(flag.Arg(0), cfg)
}
