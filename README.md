[![Build Status](https://github.com/tischda/nosleep-client/actions/workflows/build.yml/badge.svg)](https://github.com/tischda/nosleep-client/actions/workflows/build.yml)
[![Test Status](https://github.com/tischda/nosleep-client/actions/workflows/test.yml/badge.svg)](https://github.com/tischda/nosleep-client/actions/workflows/test.yml)
[![Coverage Status](https://coveralls.io/repos/tischda/nosleep-client/badge.svg)](https://coveralls.io/r/tischda/nosleep-client)
[![Linter Status](https://github.com/tischda/nosleep-client/actions/workflows/linter.yml/badge.svg)](https://github.com/tischda/nosleep-client/actions/workflows/linter.yml)
[![License](https://img.shields.io/github/license/tischda/nosleep-client.svg)](/LICENSE)
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
Usage: nosleep-client [OPTIONS] <COMMAND>

Calls the NoSleep RPC server on ADDRESS:PORT (default: 127.0.0.1:9001).
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

## Multiple processes

Another way to control the server is by registering/unregistering processes.
The server will automatically shut down when the last process is unregistered.

### Client side

Here is an example of registering and unregistering a process:

~~~
❯ nosleep-client --port 9001 register --pid 123
Connecting to RPC server at 127.0.0.1:9001 (tcp) ...
Successfully sent Register RPC

❯ nosleep-client --port 9001 read
Connecting to RPC server at 127.0.0.1:9001 (tcp) ...
Successfully sent Read RPC
Previous ThreadExecutionState flags: 0x80000000
Registered processes: [123]

❯ nosleep-client --port 9001 unregister --pid 123
Connecting to RPC server at 127.0.0.1:9001 (tcp) ...
Successfully sent Unregister RPC
~~~

### Server side

This is what it looks like server-side:

~~~
❯ nosleep-server.exe
nosleep-server starting...
ExecStateManager.System — Forcing system ON
RPC server listening on 127.0.0.1:9001 (tcp)
ExecStateManager.Register — Register process: 123
ExecStateManager.Read — Returning previous flags
ExecStateManager.Unregister — Unregister process: 123
ExecStateManager.Unregister — All processes unregistered
ExecStateManager.Shutdown - Shutting down RPC server
RPC server shutdown complete.
ThreadExecutionState cleared.
~~~

The purpose of this is to coordoniate multiple jobs that run in parallel. Without this,
the first backup process to finish would reset the execution state and potentially have
the computer going to sleep before the second job finishes. This way, we make sure that
only when all processes are done (and unregistered) the server will shut down.

## References

* [tischda/nosleep-client](/tischda/nosleep-client)
