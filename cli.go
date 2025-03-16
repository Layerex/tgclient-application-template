package main

import (
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/adrg/xdg"
)

const helpMessage = `usage: %s [-h] [--dont-save-session] [--app-id APP_ID] [--app-hash APP_HASH]

<description>

options:
  -h, --help            Show this help message and exit
  --dont-save-session   Don't save session file (and don't use already saved one)
  --app-id APP_ID       Test credentials are used by default
  --app-hash APP_HASH   Test credentials are used by default

Session file is saved to %s

<epilog>
`

type Args struct {
	DontSaveSession bool
	AppID           int32
	AppHash         string
}

func ParseArgs() Args {
	var args Args
	end := len(os.Args) - 1
	for i := 1; i < len(os.Args); i++ {
		nextArg := func() {
			if i == end {
				panic(fmt.Sprintf("Option %s requires a value", os.Args[i]))
			}
			i++
		}
		switch os.Args[i] {
		case "--app-id":
			nextArg()
			argument, err := strconv.Atoi(os.Args[i])
			if err != nil {
				panic("--app-id value has to be a 32-bit integer")
			}
			args.AppID = int32(argument)
		case "--app-hash":
			nextArg()
			if len(os.Args[i]) != 32 || !IsLowercaseHex(os.Args[i]) {
				panic("--app-hash value has to be a lowercase hex string of 32 characters")
			}
			args.AppHash = os.Args[i]
		case "--dont-save-session":
			args.DontSaveSession = true
		case "-h", "--help":
			fmt.Printf(helpMessage, os.Args[0], path.Join(xdg.DataHome, sessionFile))
			os.Exit(0)
		default:
			panic(fmt.Sprintf("Unexpected argument: %s", os.Args[i]))
		}
	}

	if args.AppID == 0 {
		if args.AppHash != "" {
			panic("--app-hash is provided, but --app-id isn't")
		}
		args.AppID = 17349
	}
	if args.AppHash == "" {
		if args.AppID == 0 {
			panic("--app-id is provided, but --app-hash isn't")
		}
		args.AppHash = "344583e45741c457fe1862106095a5eb"
	}
	return args
}
