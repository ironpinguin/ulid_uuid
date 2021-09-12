package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	getopt "github.com/pborman/getopt/v2"
)

func main() {
	var newline string = "\n"
	var err error
	var result string

	getopt.HelpColumn = 50
	getopt.DisplayWidth = 140

	fs := getopt.New()
	var (
		noNewline = fs.BoolLong("newline", 'n', "remove newline in the output")
		help      = fs.BoolLong("help", 'h', "print this help text")
	)

	fs.SetParameters("[UUID|GUID|ULID]")

	if err = fs.Getopt(os.Args, nil); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	if *noNewline {
		newline = ""
	}

	if *help {
		fs.PrintUsage(os.Stderr)
		os.Exit(0)
	}

	if fs.NArgs() != 1 {
		_, _ = fmt.Fprint(os.Stderr, "Please give one Parameter to convert!\n")
		fs.PrintUsage(os.Stderr)
		os.Exit(1)
	}

	value := fs.Arg(0)

	if result, err = checkType(value); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	_, _ = fmt.Fprintf(os.Stdout, "%s%s", result, newline)
}

func checkType(ulidUUID string) (string, error) {
	var err error

	es1, es2, converted := "", "", ""

	if converted, err = toUlid(ulidUUID); err == nil {
		return converted, nil
	}
	es1 = err.Error()

	if converted, err = toUUID(ulidUUID); err == nil {
		return converted, nil
	}
	es2 = err.Error()

	return "", errors.New("not valid ULID|UUID|GUID\n" + es1 + "\n" + es2)
}

func toUlid(uuidString string) (string, error) {
	var uuidID uuid.UUID
	var err error
	var uuidBinary []byte
	var ulidID ulid.ULID

	if uuidID, err = uuid.Parse(uuidString); err != nil {
		return "", err
	}
	uuidBinary, _ = uuidID.MarshalBinary()
	_ = ulidID.UnmarshalBinary(uuidBinary)

	return ulidID.String(), nil
}

func toUUID(ulidString string) (string, error) {
	var ulidID ulid.ULID
	var uuidID uuid.UUID
	var err error
	var ulidBinary []byte

	if ulidID, err = ulid.Parse(ulidString); err != nil {
		return "", err
	}

	ulidBinary, _ = ulidID.MarshalBinary()
	_ = uuidID.UnmarshalBinary(ulidBinary)

	return uuidID.String(), nil
}
