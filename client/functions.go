package client

import (
	"fmt"
	"time"

	"github.com/Tnze/go-mc/chat"
	"github.com/Vertixx01/wannabeflipper/flipper"
	"github.com/Vertixx01/wannabeflipper/utils"
)

func SendChatMessage(message string, delay int) error {
	time.Sleep(time.Duration(delay) * time.Second)
	err := chatHandler.SendMessage(message)
	if err != nil {
		utils.NewLogger().Error.Println(err)
	}
	return nil
}

func onGameStart() error {
	utils.NewLogger().Info.Println(fmt.Sprintf("%s Game started", utils.TimeStamp()))
	flipper.LbinFinder()
	return nil
}

func onSystemMsg(msg chat.Message, overlay bool) error {
	utils.NewLogger().Info.Println(fmt.Sprintf("%s System: %v", utils.TimeStamp(), msg))
	return nil
}

func onPlayerMsg(msg chat.Message, validated bool) error {
	utils.NewLogger().Info.Println(fmt.Sprintf("%s Player: %v", utils.TimeStamp(), msg))
	return nil
}

func onDisguisedMsg(msg chat.Message) error {
	utils.NewLogger().Info.Println(fmt.Sprintf("%s Disguised: %v", utils.TimeStamp(), msg))
	return nil
}
