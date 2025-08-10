# nosleep-client

Windows CLI utility (client) that prevents the computer from entering sleep.

The client communicates with the server via RPC and is mainly used to
shutdown the server after a task (such as a backup) has been completed.

### Install

There are no dependencies.

~~~
go install github.com/tischda/nosleep-client@latest
~~~

### Usage

~~~
Usage: nosleep-client [--port <port>] <COMMAND> | --version | --help

Calls the NoSleep RPC server on 127.0.0.1:9001.
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

EXAMPLES:

  nosleep-client --port 9015 display

  will set ThreadExecutionState to (ES_CONTINUOUS | ES_SYSTEM_REQUIRED | ES_DISPLAY_REQUIRED)
~~~

You can test the result like this (requires admin rights):

~~~
❯ powercfg -requests
DISPLAY:
None.

SYSTEM:
[PROCESS] \Device\HarddiskVolume5\src\go\nosleep-server\nosleep-server.exe

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

* [tischda/nosleep-server](/tischda/nosleep-server)
