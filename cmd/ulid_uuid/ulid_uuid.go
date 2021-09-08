package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	getopt "github.com/pborman/getopt"
	"os"
)

const (
	defaultms = "Mon Jan 02 15:04:05.999 MST 2006"
	rfc3339ms = "2006-01-02T15:04:05.999MST"
)

func main() {

	getopt.HelpColumn = 50
	getopt.DisplayWidth = 140

	fs := getopt.New()
	var (
		uuidString = fs.StringLong("uuid", 'u', "", "input of uuid to convert to ulid", "<uuidString>")
		ulidString = fs.StringLong("ulid", 'l', "", "input of ulid to convert to uuid.", "<ulidString>")
		help        = fs.BoolLong("help", 'h', "print this help text")
	)

	if err := fs.Getopt(os.Args, nil); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	if len(*uuidString) > 0 {
		_, _ = fmt.Fprintf(os.Stdout, "%s\n", toUlid(*uuidString))
		os.Exit(0)
	}

	if len(*ulidString) > 0 {
		_, _ = fmt.Fprintf(os.Stdout, "%s\n", toUuid(*ulidString))
		os.Exit(0)
	}

	if *help {
		fs.PrintUsage(os.Stderr)
		os.Exit(0)
	}

	_, _ = fmt.Fprint(os.Stdout, "Please give uuid or ulid to convert!!\n")
	fs.PrintUsage(os.Stderr)
	os.Exit(1)
}

func toUlid(uuidString string) string {
	var uuidId uuid.UUID
	var err error
	var uuidBinary []byte
	var ulidId ulid.ULID

	if uuidId, err = uuid.Parse(uuidString); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	if uuidBinary, err = uuidId.MarshalBinary(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	if err = ulidId.UnmarshalBinary(uuidBinary); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	return ulidId.String()
}

func toUuid(ulidString string) string {
	var ulidId ulid.ULID
	var uuidId uuid.UUID
	var err error
	var ulidBinary []byte

	if ulidId, err = ulid.Parse(ulidString); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	if ulidBinary, err = ulidId.MarshalBinary(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	if err = uuidId.UnmarshalBinary(ulidBinary); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	return uuidId.String()
}
