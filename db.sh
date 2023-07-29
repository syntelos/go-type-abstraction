#!/bin/bash

target=/usr/local/src/go-types-review/src/cmd/go/main.go
object=list
source=main/main.go

function usage {
    cat<<EOF>&2
Synopsis

    $0 [srv|cli] <portnum>

Description

    Start a server, then a client.  Undetached, in two
    terminals.  The server will exit when the client exits.

    The port number should be greater than 1000 and less
    than 65536 for nominal use.  The server and client
    connect on a shared port number.  Virtually any shared,
    real number will suffice as previously unoccupied port
    in the INET domain of the OS kernel IP space.

Synopsis

    $0 [list|stop]

Description

    Housekeeping belt and braces.  List or stop all 'dlv'
    processes.  Normally unnecessary.

EOF
    exit 1
}

if [ -n "${1}" ]
then
    cmd=${1}
    case "${cmd}" in
	srv|cli)
	    if [ -n "${2}" ]&&[ 0 -lt "${2}" ]&&[ 65535 -gt "${2}" ]
	    then
		port=${2}
		case "${cmd}" in
		    srv)
			dlv debug ${target} --headless --listen=127.0.0.1:${port} -- build -o ${object} ${source}
			;;
		    cli)
			dlv connect 127.0.0.1:${port}  --init bp.txt
			;;
		esac
	    else
		usage
	    fi
	;;
	list|stop)
	    if pidlist=$(ps -C dlv -o pid | egrep -v PID | awk '{print $1}') && [ -n "${pidlist}" ]
	    then
		case "${cmd}" in
		    list)
			echo ${pidlist}
			;;
		    stop)
			kill -9 ${pidlist}
			;;
		esac
	    fi
	    exit 0
	;;
    esac

else
    usage
fi
