package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"github.com/pborman/getopt"
)

func main() {
	exitCode := mainControl(os.Args)
	os.Exit(exitCode)
}

func mainControl(args []string) int {
	getopt.HelpColumn = 50
	getopt.DisplayWidth = 140

	fs := getopt.New()

	var (
		noNewline        = fs.BoolLong("newline", 'n', "remove newline in the output")
		help             = fs.BoolLong("help", 'h', "print this help text")
		newline   string = "\n"
		err       error
		result    string
	)

	fs.SetParameters("[UUID|GUID|ULID]")

	if err = fs.Getopt(args, nil); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		return 1
	}

	if *noNewline {
		newline = ""
	}

	if *help {
		fs.PrintUsage(os.Stderr)
		return 0
	}

	if fs.NArgs() != 1 {
		_, _ = fmt.Fprint(os.Stderr, "Please give one Parameter to convert!\n")
		fs.PrintUsage(os.Stderr)
		return 1
	}

	value := fs.Arg(0)

	if result, err = checkType(value); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 1
	}
	_, _ = fmt.Fprintf(os.Stdout, "%s%s", result, newline)

	return 0
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

	return "", errors.New("not valid ULID|UUID|GUID\nParsing error UUID: " + es1 + "\nParsing errro ULID: " + es2)
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
