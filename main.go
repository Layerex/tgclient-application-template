package main

import (
	"errors"
	"fmt"

	"github.com/3bl3gamer/tgclient/mtproto"
	"github.com/adrg/xdg"
)

const programName = "tgclient-application-template"
const sessionFile = programName + "/tg.session"

func (t *Telegram) GetUser() (mtproto.TL_user, error) {
	tl := t.Request(mtproto.TL_users_getUsers{ID: []mtproto.TL{mtproto.TL_inputUserSelf{}}})
	user, ok := tl.(mtproto.VectorObject)[0].(mtproto.TL_user)
	if !ok {
		return user, errors.New("TL_users_getUsers failed")
	}

	return user, nil
}

func main() {
	args := ParseArgs()

	var sessionFilePath string
	if !args.DontSaveSession {
		var err error
		sessionFilePath, err = xdg.DataFile(sessionFile)
		if err != nil {
			panic(err)
		}
	}

	var t Telegram
	err := t.SignIn(args.AppID, args.AppHash, sessionFilePath)
	if err != nil {
		panic(err)
	}

	user, err := t.GetUser()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Current user:\n%s\n", StructToString(user))
}
