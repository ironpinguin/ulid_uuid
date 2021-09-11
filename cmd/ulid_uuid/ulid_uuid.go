package main

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	getopt "github.com/pborman/getopt"
)

func main() {
	var newline string = "\n"
	var err error
	var result string

	getopt.HelpColumn = 50
	getopt.DisplayWidth = 140

	fs := getopt.New()
	var (
		uuidString = fs.StringLong("uuid", 'u', "", "input of uuid to convert to ulid", "<uuidString>")
		ulidString = fs.StringLong("ulid", 'l', "", "input of ulid to convert to uuid.", "<ulidString>")
		noNewline  = fs.BoolLong("newline", 'n', "remove newline in the output")
		help       = fs.BoolLong("help", 'h', "print this help text")
	)

	if err = fs.Getopt(os.Args, nil); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	if *noNewline {
		newline = ""
	}

	if len(*uuidString) > 0 {
		if result, err = toUlid(*uuidString); err != nil {
			os.Exit(1)
		}
		_, _ = fmt.Fprintf(os.Stdout, "%s%s", result, newline)
		os.Exit(0)
	}

	if len(*ulidString) > 0 {
		if result, err = toUuid(*uuidString); err != nil {
			os.Exit(1)
		}
		_, _ = fmt.Fprintf(os.Stdout, "%s%s", result, newline)
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

func toUlid(uuidString string) (string, error) {
	var uuidId uuid.UUID
	var err error
	var uuidBinary []byte
	var ulidId ulid.ULID

	if uuidId, err = uuid.Parse(uuidString); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		return "", err
	}

	if uuidBinary, err = uuidId.MarshalBinary(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		return "", err
	}

	if err = ulidId.UnmarshalBinary(uuidBinary); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		return "", err
	}

	return ulidId.String(), nil
}

func toUuid(ulidString string) (string, error) {
	var ulidId ulid.ULID
	var uuidId uuid.UUID
	var err error
	var ulidBinary []byte

	if ulidId, err = ulid.Parse(ulidString); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		return "", err
	}

	if ulidBinary, err = ulidId.MarshalBinary(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		return "", err
	}

	if err = uuidId.UnmarshalBinary(ulidBinary); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		return "", err
	}

	return uuidId.String(), nil
}
