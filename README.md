# waituntil

`waituntil` is a CLI command to run command repeatedly until output string is matched with string you specified.
This tool is useful to wait async command like AWS CLI.

## Getting Started

#### Download

Download binary from [release pages](https://github.com/kangaechu/waituntil/releases).

#### Run

Show help:

```shell-session
$ ./waituntil
Usage:
  waituntil [OPTIONS]

Application Options:
  -c, --command=  Check command that this tool repeatedly isseud.
  -i, --interval= Check command retry interval (sec) default:60
  -r, --retries=  Check command number of retries default:20
  -o, --ok=       Text recognized as OK (normal end)
  -n, --ng=       Text recognized as NG (with error)

Help Options:
  -h, --help      Show this help message
```

Then run.

```shell-session
$ ./waituntil --command date --interval 1 --retries 10 --ok "8" --ng "16"
2019/03/14 11:55:56 Options:
2019/03/14 11:55:56   Command:    date
2019/03/14 11:55:56   Interval:   1
2019/03/14 11:55:56   Retries:    10
2019/03/14 11:55:56   OK Message: 8
2019/03/14 11:55:56   NG Message: 16
2019/03/14 11:55:56
2019/03/14 11:55:56 output:  Thu Mar 14 11:55:56 JST 2019
2019/03/14 11:55:57 output:  Thu Mar 14 11:55:57 JST 2019
2019/03/14 11:55:58 output:  Thu Mar 14 11:55:58 JST 2019

$ echo $?
0
```

## Options

### `-c`, `--command`

Specify command to run repeatedly.

 
### `-i`, `--interval`

Specify interval time (sec.) between commands you specified option in `--command`.

### `-r`, `--retries`

Specify number of retries.
If number of retries exceeds specified retry count, waituntil returns 1 and exit.

### `-o`, `--ok` 

Specify string to be expected to the output of command.
If ok string matches with command output, waituntil returns 0 and exit.

### `-n`, `--ng` 

Specify string to not be expected to the output of command.
If ok string matches with command output, waituntil returns 1 and exit.


## Build and Run

### Prerequisites

- go 1.11+

https://golang.org/doc/install


### Installing

Get source files.

```
$ go get github.com/kangaechu/waituntil
```

Run using `go run`.

```
$ go run cmd/waituntil/main.go -c date -i 1 -r 10 -o "8" -n "16"
```

Or build and run.

```
$ go build ./cmd/waituntil
$ ./waituntil -c date -i 1 -r 10 -o "8" -n "16"
```

## Running the tests

```golang
go test -v ./pkg/waituntil
```

## License

MIT

