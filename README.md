# Command line tool to convert ULID to and from UUID

## Background
The ULID (Universally Unique Lexicographically Sortable Identifier) is a new format for unique identify.
The definition can be found under [ULID spec](https://github.com/ulid/spec)

The older and lesser readable UUID (Universally Unique Identifier) is a more used unique identifer in multiple implementation like Databases and SDK's.

Both Unique identifier are convertable from one to the other.

The idea behind this tools is to give a simple console command to convert between UUID and ULID in a simple way.

## Install
```shell
go get github.com/ironpinguin/ulid_to_from_uuid/v1/cmd/ulid_uuid
```
or get the latest Release from the GitHub Release Page
[ulid_to_from_uuid Releases](https://github.com/ironpinguin/ulid_to_from_uuid/releases)

## Command line Tool

Usage:
```text
Usage: ulid_uuid [-hn] [-l <ulidString>] [-u <uuidString>] [parameters ...]
 -h, --help               print this help text
 -l, --ulid=<ulidString>  input of ulid to convert to uuid.
 -n, --newline            remove newline in the output
 -u, --uuid=<uuidString>  input of uuid to convert to ulid
```

Examples:
```shell
$ ulid_uuid -u b5a42ffc-10b7-11ec-9135-53a1467b876b
5NMGQZR45Q27P92DAKM537Q1VB
$ ulid_uuid -l 01FF2ZVP6PPRPWQAZ5S8Q6VA7Z
017bc5fd-d8d6-b62d-cbab-e5ca2e6da8ff
$ # output without newline:
$ ulid_uuid -n -u b5a42ffc-10b7-11ec-9135-53a1467b876b
5NMGQZR45Q27P92DAKM537Q1VB
$ ulid_uuid -n -l 01FF2ZVP6PPRPWQAZ5S8Q6VA7Z
017bc5fd-d8d6-b62d-cbab-e5ca2e6da8ff
```