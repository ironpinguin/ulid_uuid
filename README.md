# Command line tool to convert ULID to and from UUID/GUID

[![Go Version](https://img.shields.io/github/go-mod/go-version/ironpinguin/ulid_uuid)](https://img.shields.io/github/go-mod/go-version/ironpinguin/ulid_uuid)
[![Coverage Status](https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/ironpinguin/97d98d096e648370e2848116f7f8289a/raw/ulid_uuid__main.json)](https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/ironpinguin/97d98d096e648370e2848116f7f8289a/raw/ulid_uuid__main.json)
[![run tests](https://github.com/ironpinguin/ulid_uuid/actions/workflows/ci.yaml/badge.svg)](https://github.com/ironpinguin/ulid_uuid/actions/workflows/ci.yaml)

## Background
The ULID (Universally Unique Lexicographically Sortable Identifier) is a new format for unique identify.
The definition can be found under [ULID spec](https://github.com/ulid/spec)

The older and lesser readable UUID/GUID ('Universally Unique Identifier'/'Globally Unique Identifier') is a more used unique identifer in multiple implementation like Databases and SDK's.

Both Unique identifier are convertable from one to the other.

The idea behind this tools is to give a simple console command to convert between UUID and ULID in a simple way.

## Install
```shell
go get github.com/ironpinguin/ulid_uuid/v1/cmd/ulid_uuid
```
or get the latest Release from the GitHub Release Page
[ulid_uuid Releases](https://github.com/ironpinguin/ulid_uuid/releases)

## Command line Tool

Usage:
```text
Usage: ulid_uuid [-hn] [UUID|GUID|ULID]
 -h, --help               print this help text
 -n, --newline            remove newline in the output
```

Examples:
```shell
$ ulid_uuid b5a42ffc-10b7-11ec-9135-53a1467b876b
5NMGQZR45Q27P92DAKM537Q1VB
$ ulid_uuid 01FF2ZVP6PPRPWQAZ5S8Q6VA7Z
017bc5fd-d8d6-b62d-cbab-e5ca2e6da8ff
$ # output without newline:
$ ulid_uuid -n b5a42ffc-10b7-11ec-9135-53a1467b876b
5NMGQZR45Q27P92DAKM537Q1VB
$ ulid_uuid -n 01FF2ZVP6PPRPWQAZ5S8Q6VA7Z
017bc5fd-d8d6-b62d-cbab-e5ca2e6da8ff
```
