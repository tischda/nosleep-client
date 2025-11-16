[![Build Status](https://github.com/tischda/nosleep-client/actions/workflows/build.yml/badge.svg)](https://github.com/tischda/nosleep-client/actions/workflows/build.yml)
[![Test Status](https://github.com/tischda/nosleep-client/actions/workflows/test.yml/badge.svg)](https://github.com/tischda/nosleep-client/actions/workflows/test.yml)
[![Coverage Status](https://coveralls.io/repos/tischda/nosleep-client/badge.svg)](https://coveralls.io/r/tischda/nosleep-client)
[![Linter Status](https://github.com/tischda/nosleep-client/actions/workflows/linter.yml/badge.svg)](https://github.com/tischda/nosleep-client/actions/workflows/linter.yml)
[![License](https://img.shields.io/github/license/tischda/nosleep-client)](/LICENSE)
[![Release](https://img.shields.io/github/release/tischda/nosleep-client.svg)](https://github.com/tischda/nosleep-client/releases/latest)


# nosleep-client

Windows CLI utility (client) that prevents the computer from entering sleep.

The client communicates with the server via RPC and is mainly used to
shutdown the server after a task (such as a backup) has been completed.

## Install

~~~
go install github.com/tischda/nosleep-client@latest
~~~

## Usage

~~~
Usage: nosleep-client [--server <address>] [--port <port>] <COMMAND>

Calls the NoSleep RPC server on SERVER:PORT (default: 127.0.0.1:9001).
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
~~~

## Examples

~~~
nosleep-client --port 9015 display
~~~

Will set ThreadExecutionState to `(ES_CONTINUOUS | ES_SYSTEM_REQUIRED | ES_DISPLAY_REQUIRED)`

You can test the result like this (requires admin rights):

~~~
❯ powercfg -requests
DISPLAY:
None.

SYSTEM:
[PROCESS] \Device\HarddiskVolume5\src\go\nosleep-client\nosleep-client.exe

AWAYMODE:
None.

EXECUTION:
None.

PERFBOOST:
None.

ACTIVELOCKSCREEN:
None.
~~~

### References

* [tischda/nosleep-client](/tischda/nosleep-client)
